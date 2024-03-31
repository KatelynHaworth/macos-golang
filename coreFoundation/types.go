package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
)

var (
	// urlPointerType represents the Golang type
	// for a url.URL structure pointer
	urlPointerType = reflect.TypeOf((*url.URL)(nil))
)

// TypeRef is a Core Foundation
// generic value pointer
type TypeRef C.CFTypeRef

// Release will instruct the kernel
// to release all memory resources
// allocated for the provided reference.
//
// If the reference is equal to 0 then
// release is not called to prevent
// crashing the thread.
func Release(ref TypeRef) {
	if ref == 0 {
		return
	}

	C.CFRelease((C.CFTypeRef)(ref))
}

// FromCFTypeRef uses reflection provided by
// the Core Foundation library to identify the
// type of information represented by the reference.
//
// If the information type is supported by this library,
// the appropriate conversion method is called to convert
// the reference to a Golang type.
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

	case C.CFURLGetTypeID():
		return FromCFURL((URLRef)(ref))

	default:
		return nil, errors.New("unsupported CoreFoundation type")
	}
}

// ToCFTypeRef will use reflection to identify the
// data type provided, based on the type provided it
// will call the appropriate conversion method to convert
// the data to a Core Foundation type reference.
func ToCFTypeRef(v interface{}) (TypeRef, error) {
	valueType := reflect.TypeOf(v)
	switch valueType.Kind() {
	case reflect.String:
		ref, err := ToCFString(v.(string))
		if err != nil {
			return 0, fmt.Errorf("convert string to CoreFoundation type: %w", err)
		}
		return (TypeRef)(ref), nil

	case reflect.Bool:
		ref, err := ToCFBoolean(v.(bool))
		if err != nil {
			return 0, fmt.Errorf("convert bool to CoreFoundation type: %w", err)
		}
		return (TypeRef)(ref), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Uint8:
		ref, err := ToCFNumber(v)
		if err != nil {
			return 0, fmt.Errorf("convert number to CoreFoundation type: %w", err)
		}
		return (TypeRef)(ref), nil

	case reflect.Slice, reflect.Array:
		if b, ok := v.([]byte); ok {
			ref, err := ToCFData(b)
			if err != nil {
				return 0, fmt.Errorf("convert byte slice to CoreFoundation type: %w", err)
			}
			return (TypeRef)(ref), nil
		} else {
			ref, err := ToCFArray(v)
			if err != nil {
				return 0, fmt.Errorf("convert slice to CoreFoundation type: %w", err)
			}
			return (TypeRef)(ref), nil
		}

	case reflect.Map:
		ref, err := ToCFDictionary(v)
		if err != nil {
			return 0, fmt.Errorf("convert map to CoreFoundation type: %w", err)
		}
		return (TypeRef)(ref), nil

	case reflect.Struct:
		if !valueType.AssignableTo(urlPointerType) {
			return 0, errors.New("unsupported Golang type")
		}

		ref, err := ToCFURL(v.(*url.URL))
		if err != nil {
			return 0, fmt.Errorf("convert URL to CoreFoundation type: %w", err)
		}
		return (TypeRef)(ref), nil

	default:
		return 0, errors.New("unsupported Golang type")
	}
}
