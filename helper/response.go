package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status"      : "Success",
		"status_code" : 200,
		"message"     : message,
		"data"		  : data,
	})
}