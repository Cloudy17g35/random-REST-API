package main

import (
	"time"
	"random-api/handlers"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

const default_timeout = 10 * time.Second

func main() {
	engine := gin.New()
	engine.GET(
	"random/mean",
	timeout.New(
	timeout.WithTimeout(default_timeout),
    timeout.WithHandler(handlers.RandomMeanHandler)),
	)

	engine.Run("localhost:8080")
}