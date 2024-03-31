package systemConfiguration

/*
#cgo CFLAGS: -mmacosx-version-min=10.6 -D__MAC_OS_X_VERSION_MAX_ALLOWED=1080
#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration

#include <SystemConfiguration/SystemConfiguration.h>

extern void goDynamicStoreCallback(SCDynamicStoreRef store, CFArrayRef changedKeys, void *context);

static SCDynamicStoreContext * CreateContext(void *info) {
	SCDynamicStoreContext * context = malloc(sizeof(SCDynamicStoreContext));
	context -> version = 0;
	context -> info = info;
	context -> retain = NULL;
	context -> release = NULL;
	context -> copyDescription = NULL;
	return context;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"

	. "github.com/LiamHaworth/macos-golang/coreFoundation"
)

// DynamicStoreCallBack is a function invoked
// when the state of a dynamic store is changed
// due to a notification.
//
// When invoked the callback will be provided with
// reference to the store that changed and an array
// of the keys that changed.
type DynamicStoreCallBack func(store *DynamicStore, changedKeys []interface{}, context interface{})

type DynamicStoreRef C.SCDynamicStoreRef

type DynamicStore struct {
	ref         C.SCDynamicStoreRef
	callback    DynamicStoreCallBack
	context     interface{}
	loopRunning bool
}

func DynamicStoreCreate(name string, callback DynamicStoreCallBack, context interface{}) (*DynamicStore, error) {
	store := new(DynamicStore)
	store.context = context
	store.callback = callback

	maskedPointer := uintptr(unsafe.Pointer(store))

	cfContext := C.CreateContext(unsafe.Pointer(&maskedPointer))
	cfName, _ := ToCFString(name)

	store.ref = C.SCDynamicStoreCreate(0, (C.CFStringRef)(cfName), (*[0]byte)(C.goDynamicStoreCallback), cfContext)

	if store.ref == 0 {
		return nil, errors.New("unable to create new dynamic store")
	}

	return store, nil
}

func (store *DynamicStore) SetNotificationKeys(keys, patterns []string) (err error) {
	cfKeys, err := ToCFArray(keys)
	if err != nil {
		return fmt.Errorf("convert keys array: %w", err)
	}

	cfPatterns, err := ToCFArray(patterns)
	if err != nil {
		return fmt.Errorf("convert patterns array: %w", err)
	}

	success := C.SCDynamicStoreSetNotificationKeys(store.ref, (C.CFArrayRef)(cfKeys), (C.CFArrayRef)(cfPatterns))
	if success == 0 {
		err = errors.New("unable to set notification keys for store")
	}

	return
}

func (store *DynamicStore) RunLoop() {
	runtime.LockOSThread()
	defer runtime.LockOSThread()

	source := C.SCDynamicStoreCreateRunLoopSource(0, store.ref, 0)
	C.CFRunLoopAddSource(C.CFRunLoopGetCurrent(), source, C.kCFRunLoopDefaultMode)

	store.loopRunning = true
	for store.loopRunning {
		C.CFRunLoopRun()
	}

	C.CFRunLoopRemoveSource(C.CFRunLoopGetCurrent(), source, C.kCFRunLoopDefaultMode)
	Release((TypeRef)(source))
}

func (store *DynamicStore) StopLoop() {
	store.loopRunning = false
	C.CFRunLoopStop(C.CFRunLoopGetCurrent())
}
