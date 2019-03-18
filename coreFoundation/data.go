package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

type DataRef C.CFDataRef

func (ref DataRef) native() C.CFDataRef {
	return (C.CFDataRef)(ref)
}

func FromCFData(ref DataRef) ([]byte, error) {
	size := C.CFDataGetLength(ref.native())
	data := make([]byte, size)

	C.CFDataGetBytes(ref.native(), C.CFRange{0, size}, (*C.UInt8)(&data[0]))

	return data, nil
}

func ToCFData(data []byte) (DataRef, error) {
	return (DataRef)(C.CFDataCreate(
		0, (*C.UInt8)(&data[0]), C.CFIndex(len(data)),
	)), nil
}
