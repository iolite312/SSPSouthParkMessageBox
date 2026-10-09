package handler

import (
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
}
