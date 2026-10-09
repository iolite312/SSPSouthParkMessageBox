package services

import (
	"ssp-sp-messaging/api/internal/core/domain"
	"ssp-sp-messaging/api/internal/core/ports"
)

type MessageService struct {
	publisher ports.MessagePublisher
}

func NewMessageService(p ports.MessagePublisher) *MessageService {
	return &MessageService{publisher: p}
}

func (s *MessageService) Send(postMessage domain.PostMessageRequest) (*domain.Message, error) {
	msg := domain.NewMessage(postMessage)
	if err := s.publisher.Publish(msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
