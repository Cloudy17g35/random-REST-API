package invalid_responses

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)


func AbortExternalServerError(
	context *gin.Context, 
	err error) {
	message := "External server error"
    context.AbortWithStatusJSON(
		http.StatusFailedDependency,
		gin.H{
		"message": message,
		"error": err,
	})
}


func AbortExternalServerBadStatus(
	context *gin.Context, 
	status_code int, 
	resp_body string) {
	message := fmt.Sprintf("Dependency error, status code from dependency site: %d", status_code)
    context.AbortWithStatusJSON(
		http.StatusFailedDependency,
		gin.H{
		"message": message,
		"error": resp_body,
	})
}
