package domain

import "github.com/google/uuid"

type Message struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type PostMessageRequest struct {
	Author string `json:"author" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

func NewMessage(postMessage PostMessageRequest) Message {
	return Message{
		ID:     uuid.NewString(),
		Author: postMessage.Author,
		Body:   postMessage.Body,
	}
}
