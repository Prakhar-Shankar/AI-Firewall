package gateway

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"ai_firewall/internal/model"
)

//ChatHandler - POST /v1/chat requests
func ChatHandler(c *gin.Context){
	var newChat model.Chat

	if err := c.ShouldBindJSON(&newChat); err != nil{
		RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	RespondSuccess(c, http.StatusCreated, "New message accepted", newChat)
}

//sends success response
func RespondSuccess(c *gin.Context, status int, message string, data interface{}){
	c.JSON(status, gin.H{
		"message": message, 
		"data": data,
	})
}

//sends error message
func RespondError(c *gin.Context, status int, errMsg string){
	c.JSON(status, gin.H{
		"error": errMsg,
	})
}