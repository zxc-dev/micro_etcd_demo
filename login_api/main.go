package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"
	"log"

	"github.com/zxc-dev/micro_etcd_demo/login_api/client"
	"github.com/zxc-dev/micro_etcd_demo/login_api/router"
)

func main() {
	// Remember to call the Init() function to initialize the go-micro client service
	reg := etcd.NewRegistry()
	service := micro.NewService(
		micro.Name("micro.api.login"),
		micro.Registry(reg),
	)
	service.Init()
	client.Init(service)

	// Start Gin Router at port 3000
	r := router.NewRouter(service)
	if err := r.Run("0.0.0.0:3000"); err != nil {
		log.Print(err.Error())
	}
}
