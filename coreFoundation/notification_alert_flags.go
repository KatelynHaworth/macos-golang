package coreFoundation

type NotificationAlertLevel uint64

const (
	UserNotificationStopAlertLevel NotificationAlertLevel = iota
	UserNotificationNoteAlertLevel
	UserNotificationCautionAlertLevel
	UserNotificationPlainAlertLevel
)
