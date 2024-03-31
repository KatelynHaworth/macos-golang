package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type StringRef C.CFStringRef

func (ref StringRef) native() C.CFStringRef {
	return (C.CFStringRef)(ref)
}

// FromCFString will convert a Core
// Foundation String reference to a
// Golang string.
//
// If the CFStringRef points to a C
// string then the pointer is converted
// using the standard libraries method
// for doing so.
//
// Otherwise the bytes of the CFString
// will be copied into a new Golang string.
//
// The CFStringRef isn't released by
// this function and will need to be
// done by the caller.
func FromCFString(ref StringRef) (s string, _ error) {
	if ref == 0 {
		return "", nil
	}

	/*
		Check if the CFStringRef is a plain C
		string pointer, if so just convert directly
		from the pointer
	*/
	if p := C.CFStringGetCStringPtr(ref.native(), C.kCFStringEncodingUTF8); p != nil {
		return C.GoString(p), nil
	}

	/*
		The CFStringRef isn't a plain CString so
		it has to be converted by copying the bytes
		and crafting a new Golang string
	*/
	length := C.CFStringGetLength(ref.native())
	if length == 0 {
		// String is already empty
		return
	}

	bufferLength := C.CFStringGetMaximumSizeForEncoding(length, C.kCFStringEncodingUTF8)
	if bufferLength == 0 {
		return "", errors.New("string is not encoded with UTF-8, unable to convert to Golang string")
	}

	var bufferFilled C.CFIndex
	buffer := make([]byte, bufferLength)

	C.CFStringGetBytes(ref.native(), C.CFRange{0, length}, C.kCFStringEncodingUTF8, C.UInt8(0), C.false, (*C.UInt8)(&buffer[0]), bufferLength, &bufferFilled)

	/*
		Make a Golang string using the byte buffer
		to prevent copying bytes and using more memory
		than needed.
	*/
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = slice.Data
	str.Len = int(bufferFilled)

	return
}

// ToCFString will create a new Core
// Foundation String reference by copying
// the Golang string data, this is to
// prevent unsafe memory releases by macOS
// libraries.
//
// As such, the CFStringRef returned must
// be freed by the caller when appropriate
// to prevent memory leaks.
func ToCFString(s string) (StringRef, error) {
	data := make([]byte, len(s))
	copy(data, s)

	return (StringRef)(C.CFStringCreateWithBytes(0, *(**C.UInt8)(unsafe.Pointer(&data)), C.CFIndex(len(s)), C.kCFStringEncodingUTF8, 0)), nil
}
