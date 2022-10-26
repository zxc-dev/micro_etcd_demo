package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zxc-dev/micro_etcd_demo/login_api/client"
)

func Login(ctx *gin.Context) {
	// example
	name := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")

	_, err := client.Login(name, passwd)
	if err != nil {
		log.Println(err.Error())
	}
	// TODO:
}
