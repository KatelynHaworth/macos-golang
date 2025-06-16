package systemConfiguration

/*
#cgo CFLAGS: -mmacosx-version-min=10.6 -D__MAC_OS_X_VERSION_MAX_ALLOWED=1080
#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration

#include <SystemConfiguration/SystemConfiguration.h>
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"

	. "github.com/LiamHaworth/macos-golang/coreFoundation"
)

//export goDynamicStoreCallback
func goDynamicStoreCallback(_ C.SCDynamicStoreRef, changedKeys C.CFArrayRef, context unsafe.Pointer) {
	array, _ := FromCFArray((ArrayRef)(changedKeys))

	handle := cgo.Handle(context)
	store, ok := handle.Value().(*DynamicStore)
	if !ok || store == nil || store.callback == nil {
		return
	}

	store.callback(store, array, store.context)
}
