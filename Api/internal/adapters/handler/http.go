package handler

import (
	"log"
	"net/http"
	"ssp-sp-messaging/api/internal/core/domain"
	"ssp-sp-messaging/api/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc ports.MessageService
}

func newHTTPHandler(svc ports.MessageService) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

func (h *HTTPHandler) postMessage(c *gin.Context) {
	var req domain.PostMessageRequest

	msg, err := h.svc.Send(req)
	if err != nil {
		log.Printf("send message failed: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "could not queue message"})
		return
	}
	c.JSON(http.StatusCreated, msg)
}
