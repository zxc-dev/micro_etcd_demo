package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zxc-dev/micro_etcd_demo/login_service/impl"
	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
	"go-micro.dev/v4"
	"log"
)

func main() {
	//impl.InitMysql()
	reg := etcd.NewRegistry()
	//defer impl.DB.Close()

	service := micro.NewService(
		micro.Name("micro.service.login"), // The service name to register in the reg
		micro.Registry(reg),
	)

	service.Init()

	loginpb.RegisterLoginHandler(service.Server(), &impl.LoginService{})

	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}
