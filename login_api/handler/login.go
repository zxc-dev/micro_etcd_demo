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

	result, err := client.Login(name, passwd)
	if err != nil {
		log.Println(err.Error())
	}
	// TODO:
	ctx.JSON(200, gin.H{
		"message": result,
	})
}
func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")

	result, err := client.Register(name, passwd)
	if err != nil {
		log.Println(err.Error())
	}
	// TODO:
	ctx.JSON(200, gin.H{
		"message": result,
	})
}
