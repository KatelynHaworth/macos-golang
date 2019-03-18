package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type NumberRef C.CFNumberRef

func (ref NumberRef) native() C.CFNumberRef {
	return (C.CFNumberRef)(ref)
}

func FromCFNumber(ref NumberRef) (interface{}, error) {
	typ := C.CFNumberGetType(ref.native())

	switch typ {
	case C.kCFNumberSInt8Type:
		var sint C.SInt8
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&sint))
		return int8(sint), nil

	case C.kCFNumberSInt16Type:
		var sint C.SInt16
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&sint))
		return int16(sint), nil

	case C.kCFNumberSInt32Type:
		var sint C.SInt32
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&sint))
		return int32(sint), nil

	case C.kCFNumberSInt64Type:
		var sint C.SInt64
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&sint))
		return int64(sint), nil

	case C.kCFNumberFloat32Type:
		var float C.Float32
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&float))
		return float32(float), nil

	case C.kCFNumberFloat64Type:
		var float C.Float64
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&float))
		return float64(float), nil

	case C.kCFNumberCharType:
		var char C.char
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&char))
		return byte(char), nil

	case C.kCFNumberShortType:
		var short C.short
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&short))
		return int16(short), nil

	case C.kCFNumberIntType:
		var i C.int
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&i))
		return int32(i), nil

	case C.kCFNumberLongType:
		var long C.long
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&long))
		return int(long), nil

	case C.kCFNumberLongLongType:
		var longlong C.longlong
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&longlong))
		return int64(longlong), nil

	case C.kCFNumberFloatType:
		var float C.float
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&float))
		return float32(float), nil

	case C.kCFNumberDoubleType:
		var double C.double
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&double))
		return float64(double), nil

	case C.kCFNumberCFIndexType:
		var index C.CFIndex
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&index))
		return int(index), nil

	case C.kCFNumberNSIntegerType:
		var nsInt C.long
		C.CFNumberGetValue(ref.native(), typ, unsafe.Pointer(&nsInt))
		return int(nsInt), nil

	default:
		return nil, errors.New("unsupported core foundation number type")
	}
}

func ToCFNumber(v interface{}) (NumberRef, error) {
	switch t := v.(type) {
	case int8:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberSInt8Type, unsafe.Pointer(&t),
		)), nil

	case int16:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberSInt16Type, unsafe.Pointer(&t),
		)), nil

	case int32:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberSInt32Type, unsafe.Pointer(&t),
		)), nil

	case int64:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberSInt64Type, unsafe.Pointer(&t),
		)), nil

	case float32:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberFloat32Type, unsafe.Pointer(&t),
		)), nil

	case float64:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberFloat64Type, unsafe.Pointer(&t),
		)), nil

	case uint8:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberCharType, unsafe.Pointer(&t),
		)), nil

	case int:
		return (NumberRef)(C.CFNumberCreate(
			0, C.kCFNumberIntType, unsafe.Pointer(&t),
		)), nil

	default:
		return 0, errors.New("unsupported Golang number type")
	}
}
