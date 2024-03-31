package systemConfiguration

/*
#cgo CFLAGS: -mmacosx-version-min=10.6 -D__MAC_OS_X_VERSION_MAX_ALLOWED=1080
#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration

#include <SystemConfiguration/SystemConfiguration.h>
*/
import "C"
import (
	"unsafe"

	. "github.com/LiamHaworth/macos-golang/coreFoundation"
)

func DynamicStoreCopyComputerName(store *DynamicStore) string {
	var (
		ref      C.SCDynamicStoreRef
		encoding uint8
	)

	if store != nil {
		ref = store.ref
	}

	value := C.SCDynamicStoreCopyComputerName(ref, (*C.uint)(unsafe.Pointer(&encoding)))
	defer Release((TypeRef)(value))

	s, _ := FromCFString((StringRef)(value))
	return s
}

func DynamicStoreCopyConsoleUser(store *DynamicStore) (string, uint, uint) {
	var (
		ref      C.SCDynamicStoreRef
		uid, gid uint
	)

	if store != nil {
		ref = store.ref
	}

	value := C.SCDynamicStoreCopyConsoleUser(ref, (*C.uint)(unsafe.Pointer(&uid)), (*C.uint)(unsafe.Pointer(&gid)))
	defer Release((TypeRef)(value))

	s, _ := FromCFString((StringRef)(value))
	return s, uid, gid
}

func DynamicStoreCopyLocalHostName(store *DynamicStore) string {
	var ref C.SCDynamicStoreRef

	if store != nil {
		ref = store.ref
	}

	value := C.SCDynamicStoreCopyLocalHostName(ref)
	defer Release((TypeRef)(value))

	s, _ := FromCFString((StringRef)(value))
	return s
}

func DynamicStoreCopyLocation(store *DynamicStore) string {
	var ref C.SCDynamicStoreRef

	if store != nil {
		ref = store.ref
	}

	value := C.SCDynamicStoreCopyLocation(ref)
	defer Release((TypeRef)(value))

	s, _ := FromCFString((StringRef)(value))
	return s
}
