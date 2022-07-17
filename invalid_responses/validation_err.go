package invalid_responses

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/gin-gonic/gin"
)


func AbortValidationError(
	context *gin.Context, 
	validation_errors url.Values) {
	fmt.Printf("validation errors: %v", validation_errors)
	message := "Validation Error, please check your inputs"
    context.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{
		"message": message,
		"error": validation_errors,
	})
}

