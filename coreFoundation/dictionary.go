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

type DictionaryRef C.CFDictionaryRef

func (ref DictionaryRef) native() C.CFDictionaryRef {
	return (C.CFDictionaryRef)(ref)
}

func FromCFDictionary(ref DictionaryRef) (map[interface{}]interface{}, error) {
	if ref == 0 {
		return map[interface{}]interface{}{}, nil
	}

	size := C.CFDictionaryGetCount(ref.native())

	dictionary := make(map[interface{}]interface{}, 0)
	keys := make([]TypeRef, size)
	values := make([]TypeRef, size)

	C.CFDictionaryGetKeysAndValues(ref.native(), (*unsafe.Pointer)(unsafe.Pointer(&keys[0])), (*unsafe.Pointer)(unsafe.Pointer(&values[0])))

	for i := 0; i < int(size); i++ {
		key, err := FromCFTypeRef(keys[i])
		if err != nil {
			return nil, fmt.Errorf("convert key from CFDictionaryRef: %w", err)
		}

		value, err := FromCFTypeRef(values[i])
		if err != nil {
			return nil, fmt.Errorf("convert value from CFDictionRef: %w", err)
		}

		dictionary[key] = value
	}

	return dictionary, nil
}

func ToCFDictionary(dictionary interface{}) (DictionaryRef, error) {
	m := reflect.ValueOf(dictionary)
	if m.Kind() != reflect.Map || m.Len() == 0 {
		return 0, nil
	}

	keys := make([]TypeRef, 0, m.Len())
	values := make([]TypeRef, 0, m.Len())

	for i := range m.MapKeys() {
		key, err := ToCFTypeRef(m.MapKeys()[i].Interface())
		if err != nil {
			return 0, fmt.Errorf("convert key for CFDictionaryRef: %w", err)
		}

		value, err := ToCFTypeRef(m.MapIndex(m.MapKeys()[i]).Interface())
		if err != nil {
			return 0, fmt.Errorf("convert value for CFDictionaryRef: %w", err)
		}

		keys = append(keys, key)
		values = append(values, value)
	}

	return (DictionaryRef)(C.CFDictionaryCreate(
		0, (*unsafe.Pointer)(unsafe.Pointer(&keys[0])), (*unsafe.Pointer)(unsafe.Pointer(&values[0])), C.CFIndex(m.Len()), nil, nil,
	)), nil
}
