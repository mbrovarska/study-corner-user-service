package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	Logger *log.Logger
}

func NewHealthHandler(logger *log.Logger) *HealthHandler {
	return &HealthHandler{Logger: logger}
}
 
func (h *HealthHandler) CheckHealth(c *gin.Context) {
	h.Logger.Println("DEBUG: Health check requested")
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
		"service": "user-service",
		"version": "v1.0.0",
		"message": "Service is operational",
	})
}