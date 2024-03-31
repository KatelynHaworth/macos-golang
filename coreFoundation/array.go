package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

// ArrayRef is a Core Foundation pointer
// to an array type
type ArrayRef C.CFArrayRef

// native casts the array reference from a
// Golang type to the native C type for use
// in method calls
func (ref ArrayRef) native() C.CFArrayRef {
	return (C.CFArrayRef)(ref)
}

// FromCFArray will convert a Core
// Foundation Array reference to a
// Golang slice.
//
// Elements of the array are also
// converted to their respective Golang
// type.
//
// The reference is not released by
// this function and must be done by
// the caller.
func FromCFArray(ref ArrayRef) ([]interface{}, error) {
	if ref == 0 {
		return []interface{}{}, nil
	}

	size := C.CFArrayGetCount((C.CFArrayRef)(ref))
	elements := make([]TypeRef, size)

	C.CFArrayGetValues((C.CFArrayRef)(ref), C.CFRange{0, size}, (*unsafe.Pointer)(unsafe.Pointer(&elements[0])))

	array := make([]interface{}, size)
	for i := range elements {
		element, err := FromCFTypeRef(elements[i])
		if err != nil {
			return nil, fmt.Errorf("convert from CFArrayRef element: %w", err)
		}

		array[i] = element
	}

	return array, nil
}

// ToCFArray will create a new Core Foundation
// Array reference from a Golang slice,
// each element of the slice is converted
// to it's respective CF type.
//
// The CFArray reference must be released
// once finished with to prevent memory leaks.
func ToCFArray(v interface{}) (ArrayRef, error) {
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Slice || s.Len() == 0 {
		return 0, nil
	}

	values := make([]TypeRef, s.Len())
	for i := 0; i < s.Len(); i++ {
		value, err := ToCFTypeRef(s.Index(i).Interface())
		if err != nil {
			return 0, fmt.Errorf("convert to CFArrayRef element: %w", err)
		}

		values[i] = value
	}

	return (ArrayRef)(C.CFArrayCreate(0, (*unsafe.Pointer)(unsafe.Pointer(&values[0])), C.CFIndex(s.Len()), &C.kCFTypeArrayCallBacks)), nil
}
