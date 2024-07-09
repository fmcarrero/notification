package domain

import "time"

type NotificationRepository interface {
	SaveNotification(notification Notification, validEmails []time.Time)
	SearchNotification(notification Notification) []time.Time
}
