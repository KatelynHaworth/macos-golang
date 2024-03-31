package coreFoundation

import "C"
import (
	"net/url"
	"time"
)

type NotificationAlertOptions struct {
	Timeout              time.Duration
	Level                NotificationAlertLevel
	Icon                 *url.URL
	Localization         *url.URL
	AlertHeader          string
	AlertMessage         string
	DefaultButtonTitle   string
	AlternateButtonTitle string
	OtherButtonTitle     string
}

func (options NotificationAlertOptions) timeout() C.double {
	return C.double(options.Timeout.Seconds())
}

func (options NotificationAlertOptions) iconURL() (URLRef, error) {
	if options.Icon == nil {
		return 0, nil
	}

	return ToCFURL(options.Icon)
}

func (options NotificationAlertOptions) localizationURL() (URLRef, error) {
	if options.Localization == nil {
		return 0, nil
	}

	return ToCFURL(options.Localization)
}

func (options NotificationAlertOptions) alertHeader() StringRef {
	if len(options.AlertHeader) == 0 {
		return 0
	}

	ref, _ := ToCFString(options.AlertHeader)
	return ref
}

func (options NotificationAlertOptions) alertMessage() StringRef {
	if len(options.AlertMessage) == 0 {
		return 0
	}

	ref, _ := ToCFString(options.AlertMessage)
	return ref
}

func (options NotificationAlertOptions) defaultButtonTitle() StringRef {
	if len(options.DefaultButtonTitle) == 0 {
		return 0
	}

	ref, _ := ToCFString(options.DefaultButtonTitle)
	return ref
}

func (options NotificationAlertOptions) alternateButtonTitle() StringRef {
	if len(options.AlternateButtonTitle) == 0 {
		return 0
	}

	ref, _ := ToCFString(options.AlternateButtonTitle)
	return ref
}

func (options NotificationAlertOptions) otherButtonTitle() StringRef {
	if len(options.OtherButtonTitle) == 0 {
		return 0
	}

	ref, _ := ToCFString(options.OtherButtonTitle)
	return ref
}
