package repository

import (
	"github.com/fmcarrero/notification/src/notifications/domain"
	"time"
)

type notificationMemoryRepository struct {
	sentEmails map[domain.Recipient]map[domain.NotificationType][]time.Time
}

func NewNotificationMemoryRepository() domain.NotificationRepository {
	return &notificationMemoryRepository{
		sentEmails: make(map[domain.Recipient]map[domain.NotificationType][]time.Time),
	}
}
func (re *notificationMemoryRepository) SaveNotification(notification domain.Notification, validEmails []time.Time) {
	now := time.Now()
	// Initialize the map if not already done
	if re.sentEmails[notification.Recipient] == nil {
		re.sentEmails[notification.Recipient] = make(map[domain.NotificationType][]time.Time)
	}

	// Record the timestamp
	re.sentEmails[notification.Recipient][notification.Type] = append(validEmails, now)
}
func (re *notificationMemoryRepository) SearchNotification(notification domain.Notification) []time.Time {
	return re.sentEmails[notification.Recipient][notification.Type]
}
