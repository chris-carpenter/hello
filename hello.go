package main

import (
	"github.com/chris-carpenter/hello/configuration"
	"github.com/chris-carpenter/hello/messages"
	"github.com/chris-carpenter/hello/health"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration.Config()
	router := gin.Default()
	router.GET("/hello", messages.Hello)
	router.GET("/headers", health.Headers)
	_ = router.Run()
}