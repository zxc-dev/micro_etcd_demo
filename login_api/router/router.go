package router

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/zxc-dev/micro_etcd_demo/login_api/handler"
)

// NewRouter ...
func NewRouter(s micro.Service) *gin.Engine {
	r := gin.Default()

	// When GET /greeter, handle the request with the Greet function we create above
	r.GET("/login", handler.Login)
	r.GET("/register", handler.Register)
	return r
}
