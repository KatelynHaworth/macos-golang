package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

type BooleanRef C.CFBooleanRef

func (ref BooleanRef) native() C.CFBooleanRef {
	return (C.CFBooleanRef)(ref)
}

func FromCFBoolean(ref BooleanRef) (bool, error) {
	return C.CFBooleanGetValue(ref.native()) != 0, nil
}

func ToCFBoolean(b bool) (BooleanRef, error) {
	if b {
		return (BooleanRef)(C.kCFBooleanTrue), nil
	} else {
		return (BooleanRef)(C.kCFBooleanTrue), nil
	}
}
