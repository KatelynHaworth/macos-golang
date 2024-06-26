package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"fmt"
	"net/url"
)

type URLRef C.CFURLRef

func (ref URLRef) native() C.CFURLRef {
	return (C.CFURLRef)(ref)
}

func FromCFURL(ref URLRef) (*url.URL, error) {
	if ref == 0 {
		return nil, nil
	}

	stringRef := (StringRef)(C.CFURLGetString(ref.native()))
	defer Release((TypeRef)(stringRef))

	urlValue, _ := FromCFString(stringRef)
	parsed, err := url.Parse(urlValue)
	if err != nil {
		return nil, fmt.Errorf("parse URL value: %w", err)
	}

	return parsed, nil
}

func ToCFURL(value *url.URL) (URLRef, error) {
	stringRef, _ := ToCFString(value.String())
	defer Release((TypeRef)(stringRef))

	return (URLRef)(C.CFURLCreateWithString(0, stringRef.native(), 0)), nil
}
