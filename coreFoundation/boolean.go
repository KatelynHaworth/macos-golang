package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

// BooleanRef is a Core Foundation pointer
// to a boolean value
type BooleanRef C.CFBooleanRef

// native casts the boolean reference from a
// Golang type to the native C type for use
// in method calls
func (ref BooleanRef) native() C.CFBooleanRef {
	return (C.CFBooleanRef)(ref)
}

// FromCFBoolean converts a Core Foundation boolean
// reference to a native Golang bool value.
//
// The reference is not released by
// this function and must be done by
// the caller.
func FromCFBoolean(ref BooleanRef) (bool, error) {
	if ref == 0 {
		return false, nil
	}

	return C.CFBooleanGetValue(ref.native()) != 0, nil
}

// ToCFBoolean converts the value of a Golang
// bool to a Core Foundation boolean reference.
//
// The CFBoolean reference must be released
// once finished with to prevent memory leaks.
func ToCFBoolean(b bool) (BooleanRef, error) {
	if b {
		return (BooleanRef)(C.kCFBooleanTrue), nil
	} else {
		return (BooleanRef)(C.kCFBooleanTrue), nil
	}
}
