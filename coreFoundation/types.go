package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"reflect"

	"github.com/pkg/errors"
)

type TypeRef C.CFTypeRef

func Release(ref TypeRef) {
	C.CFRelease((C.CFTypeRef)(ref))
}

func FromCFTypeRef(ref TypeRef) (interface{}, error) {
	switch C.CFGetTypeID((C.CFTypeRef)(ref)) {
	case C.CFStringGetTypeID():
		return FromCFString((StringRef)(ref))

	case C.CFBooleanGetTypeID():
		return FromCFBoolean((BooleanRef)(ref))

	case C.CFNumberGetTypeID():
		return FromCFNumber((NumberRef)(ref))

	case C.CFDataGetTypeID():
		return FromCFData((DataRef)(ref))

	case C.CFArrayGetTypeID():
		return FromCFArray((ArrayRef)(ref))

	case C.CFDictionaryGetTypeID():
		return FromCFDictionary((DictionaryRef)(ref))

	default:
		return nil, errors.New("unsupported CoreFoundation type")
	}
}

func ToCFTypeRef(v interface{}) (TypeRef, error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		ref, err := ToCFString(v.(string))
		if err != nil {
			return 0, errors.Wrap(err, "convert string to CoreFoundation type")
		}
		return (TypeRef)(ref), nil

	case reflect.Bool:
		ref, err := ToCFBoolean(v.(bool))
		if err != nil {
			return 0, errors.Wrap(err, "convert bool to CoreFoundation type")
		}
		return (TypeRef)(ref), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Uint8:
		ref, err := ToCFNumber(v)
		if err != nil {
			return 0, errors.Wrap(err, "convert number to CoreFoundation type")
		}
		return (TypeRef)(ref), nil

	case reflect.Slice, reflect.Array:
		if b, ok := v.([]byte); ok {
			ref, err := ToCFData(b)
			if err != nil {
				return 0, errors.Wrap(err, "convert byte slice to CoreFoundation type")
			}
			return (TypeRef)(ref), nil
		} else {
			ref, err := ToCFArray(v)
			if err != nil {
				return 0, errors.Wrap(err, "convert slice to CoreFoundation type")
			}
			return (TypeRef)(ref), nil
		}

	case reflect.Map:
		ref, err := ToCFDictionary(v)
		if err != nil {
			return 0, errors.Wrap(err, "convert map to CoreFoundation type")
		}
		return (TypeRef)(ref), nil

	default:
		return 0, errors.New("unsupported Golang type")
	}
}
