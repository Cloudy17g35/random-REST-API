package external_requests

import (
	"github.com/gin-gonic/gin"
)


func QueryRequest(context *gin.Context) (string, string) {
	user_request := context.Request.URL.Query()
	l := user_request.Get("length")
	rq := user_request.Get("requests")
	return rq, l
}