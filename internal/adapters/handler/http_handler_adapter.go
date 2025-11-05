package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// HTTPHandler is the adapter for the Gin web framework
type HTTPHandler struct {
	Logger *log.Logger
	HealthHandler *HealthHandler
}

//NewHTTPHandler is the Fx provider for the HTTPHandler struct
func NewHTTPHandler(logger *log.Logger, health *HealthHandler) *HTTPHandler {
	return &HTTPHandler{
		Logger:	logger,
		HealthHandler: health,
	}
}

func(h *HTTPHandler) RegisterRoutes(router *gin.Engine) {
	// Register the Health Check handler globally (outside the API group)
	router.GET("/health", h.HealthHandler.CheckHealth)
}