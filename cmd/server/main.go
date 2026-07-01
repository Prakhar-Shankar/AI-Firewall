package main 

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Chat struct{
	Chatvalue string `json:"request" binding:"required"`
}

func SetupRouter() *gin.Engine{
	r := gin.Default()

	r.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"Status":"Server is up and running"})
	})

	r.POST("/v1/chat", func(c *gin.Context){
		var newChat Chat

		if err := c.ShouldBindJSON(&newChat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "New message accepted",
			"data": newChat,
		})
	})
	return r
}

func main(){
	r:= SetupRouter()

	r.Run(":3300")
}