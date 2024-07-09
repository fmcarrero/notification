package domain

import "time"

type NotificationType string

const (
	Status    NotificationType = "status"
	News      NotificationType = "news"
	Marketing NotificationType = "marketing"
)

type RateLimitRule struct {
	MaxEmails int
	Duration  time.Duration
}

type Recipient struct {
	Email string
}

type Notification struct {
	Type      NotificationType
	Recipient Recipient
	Message   string
}

var RateLimitRules = map[NotificationType]RateLimitRule{
	Status:    {MaxEmails: 2, Duration: time.Minute},
	News:      {MaxEmails: 1, Duration: 24 * time.Hour},
	Marketing: {MaxEmails: 3, Duration: time.Hour},
}
