package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) SaveNotification(notification Notification, validEmails []time.Time) {
	m.Called(notification, validEmails)
}

func (m *MockNotificationRepository) SearchNotification(notification Notification) []time.Time {
	args := m.Called(notification)
	return args.Get(0).([]time.Time)
}

func TestSendFirstNotification(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	rateLimitRules := map[NotificationType]RateLimitRule{
		Status: {MaxEmails: 2, Duration: time.Minute},
		News:   {MaxEmails: 1, Duration: 24 * time.Hour},
	}

	service := NewNotificationService(mockRepo, rateLimitRules)
	recipient := Recipient{Email: "user@example.com"}
	notification := Notification{
		Type:      Status,
		Recipient: recipient,
		Message:   "Test Status Update",
	}

	mockRepo.On("SearchNotification", notification).Return([]time.Time{})
	mockRepo.On("SaveNotification", notification, mock.Anything).Return()

	err := service.SendNotification(notification)
	assert.NoError(t, err, "Expected no error on first send")
	mockRepo.AssertCalled(t, "SaveNotification", notification, mock.Anything)
}

func TestSendSecondNotificationWithinLimit(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	rateLimitRules := map[NotificationType]RateLimitRule{
		Status: {MaxEmails: 2, Duration: time.Minute},
		News:   {MaxEmails: 1, Duration: 24 * time.Hour},
	}

	service := NewNotificationService(mockRepo, rateLimitRules)
	recipient := Recipient{Email: "user@example.com"}
	notification := Notification{
		Type:      Status,
		Recipient: recipient,
		Message:   "Test Status Update",
	}

	now := time.Now()
	mockRepo.On("SearchNotification", notification).Return([]time.Time{now})
	mockRepo.On("SaveNotification", notification, mock.Anything).Return()

	err := service.SendNotification(notification)
	assert.NoError(t, err, "Expected no error on second send")
	mockRepo.AssertCalled(t, "SaveNotification", notification, mock.Anything)
}

func TestExceedRateLimit(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	rateLimitRules := map[NotificationType]RateLimitRule{
		Status: {MaxEmails: 2, Duration: time.Minute},
		News:   {MaxEmails: 1, Duration: 24 * time.Hour},
	}

	service := NewNotificationService(mockRepo, rateLimitRules)
	recipient := Recipient{Email: "user@example.com"}
	notification := Notification{
		Type:      Status,
		Recipient: recipient,
		Message:   "Test Status Update",
	}

	now := time.Now()
	mockRepo.On("SearchNotification", notification).Return([]time.Time{now, now.Add(30 * time.Second)})

	err := service.SendNotification(notification)
	assert.Error(t, err, "Expected rate limit exceeded error")
	mockRepo.AssertNotCalled(t, "SaveNotification", notification, mock.Anything)
}
