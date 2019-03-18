package systemConfiguration

/*
#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
#include <SystemConfiguration/SystemConfiguration.h>

static CFStringRef DynamicStoreKeyCreate(CFAllocatorRef allocator, CFStringRef fmt, void *args[]) {
	return SCDynamicStoreKeyCreate(allocator, fmt, args);
}
*/
import "C"
import (
	"unsafe"

	. "github.com/LiamHaworth/macos-golang/coreFoundation"
)

const (
	DynamicStoreDomainFile   = "File:"
	DynamicStoreDomainPlugin = "Plugin:"
	DynamicStoreDomainSetup  = "Setup:"
	DynamicStoreDomainState  = "State:"
	DynamicStoreDomainPrefs  = "Prefs:"
)

func DynamicStoreKeyCreate(format string, args ...interface{}) string {
	fmt, _ := ToCFString(format)
	defer Release((TypeRef)(fmt))

	cArgs := make([]unsafe.Pointer, len(args))
	for i := range args {
		value, _ := ToCFTypeRef(args[i])
		cArgs[i] = (unsafe.Pointer)(uintptr(value))
	}

	key := C.DynamicStoreKeyCreate(0, (C.CFStringRef)(fmt), (*unsafe.Pointer)(cArgs[0]))
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateNetworkGlobalEntity(domain, entity string) string {
	dmn, _ := ToCFString(domain)
	defer Release((TypeRef)(dmn))

	ent, _ := ToCFString(entity)
	defer Release((TypeRef)(ent))

	key := C.SCDynamicStoreKeyCreateNetworkGlobalEntity(0, (C.CFStringRef)(dmn), (C.CFStringRef)(ent))
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateNetworkInterface(domain string) string {
	dmn, _ := ToCFString(domain)
	defer Release((TypeRef)(dmn))

	key := C.SCDynamicStoreKeyCreateNetworkInterface(0, (C.CFStringRef)(dmn))
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateNetworkInterfaceEntity(domain, ifname, entity string) string {
	dmn, _ := ToCFString(domain)
	defer Release((TypeRef)(dmn))

	intf, _ := ToCFString(ifname)
	defer Release((TypeRef)(intf))

	ent, _ := ToCFString(entity)
	defer Release((TypeRef)(ent))

	key := C.SCDynamicStoreKeyCreateNetworkInterfaceEntity(0, (C.CFStringRef)(dmn), (C.CFStringRef)(intf), (C.CFStringRef)(ent))
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateNetworkServiceEntity(domain, serviceId, entity string) string {
	dmn, _ := ToCFString(domain)
	defer Release((TypeRef)(dmn))

	id, _ := ToCFString(serviceId)
	defer Release((TypeRef)(id))

	ent, _ := ToCFString(entity)
	defer Release((TypeRef)(ent))

	key := C.SCDynamicStoreKeyCreateNetworkServiceEntity(0, (C.CFStringRef)(dmn), (C.CFStringRef)(id), (C.CFStringRef)(ent))
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateComputerName() string {
	key := C.SCDynamicStoreKeyCreateComputerName(0)
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateConsoleUser() string {
	key := C.SCDynamicStoreKeyCreateConsoleUser(0)
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateHostNames() string {
	key := C.SCDynamicStoreKeyCreateHostNames(0)
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateLocation() string {
	key := C.SCDynamicStoreKeyCreateLocation(0)
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}

func DynamicStoreKeyCreateProxies() string {
	key := C.SCDynamicStoreKeyCreateProxies(0)
	defer Release((TypeRef)(key))

	s, _ := FromCFString((StringRef)(key))
	return s
}
