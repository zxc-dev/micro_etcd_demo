package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zxc-dev/micro_etcd_demo/login_api/handler"
	"go-micro.dev/v4"
)

// NewRouter ...
func NewRouter(s micro.Service) *gin.Engine {
	r := gin.Default()

	// When GET /greeter, handle the request with the Greet function we create above
	r.GET("/login", handler.Login)
	return r
}
