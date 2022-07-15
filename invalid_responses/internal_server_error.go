package invalid_responses

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func AbortInternalServerError(context *gin.Context) {
	message := "Internal server error"
    context.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{
		"message": message,
	})
}