package ports

import (
	"ssp-sp-messaging/api/internal/core/domain"
)

type MessageService interface {
	Send(postMessage domain.PostMessageRequest) (*domain.Message, error)
}

type MessagePublisher interface {
	Publish(msg domain.Message) error
}
