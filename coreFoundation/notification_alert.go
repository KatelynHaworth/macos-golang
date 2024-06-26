package coreFoundation

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func UserNotificationDisplayAlert(options NotificationAlertOptions) (response uint64, _ error) {
	icon, err := options.iconURL()
	if err != nil {
		return 0, fmt.Errorf("convert icon URL to native: %w", err)
	} else if icon > 0 {
		defer Release((TypeRef)(icon))
	}

	localization, err := options.localizationURL()
	if err != nil {
		return 0, fmt.Errorf("convert localization URL to native: %w", err)
	} else if localization > 0 {
		defer Release((TypeRef)(localization))
	}

	C.CFUserNotificationDisplayAlert(
		options.timeout(), C.ulong(options.Level), icon.native(), 0, localization.native(), options.alertHeader().native(), options.alertMessage().native(),
		options.defaultButtonTitle().native(), options.alternateButtonTitle().native(), options.otherButtonTitle().native(), (*C.CFOptionFlags)(unsafe.Pointer(&response)),
	)

	return
}
