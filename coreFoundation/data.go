package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

// DataRef is a Core Foundation pointer
// to an array of bytes in memory
type DataRef C.CFDataRef

// native casts the data reference from a
// Golang type to the native C type for use
// in method calls
func (ref DataRef) native() C.CFDataRef {
	return (C.CFDataRef)(ref)
}

// FromCFData will convert a Core Foundation
// byte array to a Golang slice containing the
// value of the byte array.
//
// The value of the CFData reference is copied
// so that the value of the slice doesn't point
// to the same location in memory, this prevents
// changes to the slice from effecting the value
// of the CFData reference.
//
// The reference is not released by
// this function and must be done by
// the caller.
func FromCFData(ref DataRef) ([]byte, error) {
	if ref == 0 {
		return []byte{}, nil
	}

	size := C.CFDataGetLength(ref.native())
	data := make([]byte, size)

	C.CFDataGetBytes(ref.native(), C.CFRange{0, size}, (*C.UInt8)(&data[0]))

	return data, nil
}

// ToCFData will create a new Core Foundation
// data reference from a Golang slice, the value
// of the slice is copied to internal memory to
// prevent changes in the slice affecting the data
// reference.
//
// The CFData reference must be released
// once finished with to prevent memory leaks.
func ToCFData(data []byte) (DataRef, error) {
	return (DataRef)(C.CFDataCreate(
		0, (*C.UInt8)(&data[0]), C.CFIndex(len(data)),
	)), nil
}
