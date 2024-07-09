package repository

import (
	"testing"
	"time"

	"github.com/fmcarrero/notification/src/notifications/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveNotification(t *testing.T) {
	repo := NewNotificationMemoryRepository()
	recipient := domain.Recipient{Email: "user@example.com"}
	notificationType := domain.Status
	notification := domain.Notification{
		Type:      notificationType,
		Recipient: recipient,
		Message:   "Test Status Update",
	}

	// Save the notification
	repo.SaveNotification(notification, []time.Time{})

	// Retrieve the notifications for the recipient
	timestamps := repo.SearchNotification(notification)

	// Check that there is exactly one timestamp
	assert.Equal(t, 1, len(timestamps), "Expected one timestamp")
}

func TestSearchNotification(t *testing.T) {
	repo := NewNotificationMemoryRepository()
	recipient := domain.Recipient{Email: "user@example.com"}
	notificationType := domain.Status
	notification := domain.Notification{
		Type:      notificationType,
		Recipient: recipient,
		Message:   "Test Status Update",
	}

	// Save the notification
	now := time.Now()
	validEmails := []time.Time{now}
	repo.SaveNotification(notification, validEmails)

	// Retrieve the notifications for the recipient
	timestamps := repo.SearchNotification(notification)

	// Check that the retrieved timestamps match the saved ones
	assert.Equal(t, 2, len(timestamps), "Expected two timestamps")
	assert.Equal(t, now, timestamps[0], "Expected the first timestamp to match")
}
