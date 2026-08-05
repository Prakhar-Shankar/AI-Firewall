package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ai_firewall/internal/gateway"
)

// SetupRouter builds and returns the configured Gin engine.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "Server is up and running"})
	})

	r.POST("/v1/chat", gateway.ChatHandler)

	return r
}