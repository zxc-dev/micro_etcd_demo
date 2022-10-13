package router

import (
	"greeter_api/client"
	"github.com/gin-gonic/gin"
)

// NewRouter ...
func NewRouter() *gin.Engine {
	route := gin.Default()

	// When GET /greeter, handle the request with the Greet function we create above
	route.GET("/greeter", client.Greet)
	return route
}
