package main

import (
	"log"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/zxc-dev/micro_etcd_demo/login_service/impl"
	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
	"go-micro.dev/v4"
)

func main() {
	reg := etcd.NewRegistry()

	service := micro.NewService(
		micro.Name("micro.service.greeter"), // The service name to register in the reg
		micro.Registry(reg),
	)

	service.Init()

	loginpb.RegisterLoginHandler(service.Server(), &impl.LoginService{})

	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}
