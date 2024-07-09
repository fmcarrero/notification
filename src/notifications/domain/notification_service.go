package domain

import (
	"fmt"
	"sync"
	"time"
)

type NotificationService struct {
	mu                     sync.Mutex
	rateLimitRules         map[NotificationType]RateLimitRule
	notificationRepository NotificationRepository
}

func NewNotificationService(notificationRepository NotificationRepository,
	rateLimitRules map[NotificationType]RateLimitRule) *NotificationService {
	return &NotificationService{
		rateLimitRules:         rateLimitRules,
		notificationRepository: notificationRepository,
	}
}

func (s *NotificationService) SendNotification(notification Notification) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	rule, exists := s.rateLimitRules[notification.Type]
	if !exists {
		return fmt.Errorf("no rate limit rule for notification type: %s", notification.Type)
	}

	now := time.Now()

	// Remove old timestamps
	validEmails := s.notificationRepository.SearchNotification(notification)
	for i := 0; i < len(validEmails); i++ {
		if validEmails[i].Add(rule.Duration).After(now) {
			validEmails = validEmails[i:]
			break
		}
	}

	// Check if the limit has been reached
	if len(validEmails) >= rule.MaxEmails {
		return fmt.Errorf("rate limit exceeded for recipient: %s and notification type: %s", notification.Recipient.Email, notification.Type)
	}

	// Send the email (here you would integrate with your email sending logic)
	fmt.Printf("Sending %s notification to %s: %s\n", notification.Type, notification.Recipient.Email, notification.Message)
	s.notificationRepository.SaveNotification(notification, validEmails)
	return nil
}
