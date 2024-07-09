package main

import (
	"fmt"
	"github.com/fmcarrero/notification/src/notifications/domain"
	"github.com/fmcarrero/notification/src/notifications/infrastructure/repository"
	"time"
)

func main() {
	service := domain.NewNotificationService(repository.NewNotificationMemoryRepository(), domain.RateLimitRules)

	notifications := []domain.Notification{
		{Type: domain.Status, Recipient: domain.Recipient{Email: "user@example.com"}, Message: "Status update"},
		{Type: domain.News, Recipient: domain.Recipient{Email: "user@example.com"}, Message: "Daily news"},
		{Type: domain.Marketing, Recipient: domain.Recipient{Email: "user@example.com"}, Message: "Special offer"},
		{Type: domain.News, Recipient: domain.Recipient{Email: "user@example.com"}, Message: "Special offer rejected"},
	}

	for _, notification := range notifications {
		err := service.SendNotification(notification)
		if err != nil {
			fmt.Println("Error:", err)
		}
		time.Sleep(2 * time.Second)
	}
}
