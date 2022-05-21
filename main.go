package main

import (
	"counter/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	var counter middleware.Counter
	counter.Init(middleware.Flush2brokerStderr)
	r.Use(counter())
	// Handle request from PMS
	// r.POST(...) ...
}
