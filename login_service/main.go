package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/zxc-dev/micro_etcd_demo/login_service/impl"
	"github.com/zxc-dev/micro_etcd_demo/login_service/model"
	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
	"log"
)

func main() {
	model.InitMysql()
	reg := etcd.NewRegistry()
	defer model.DB.Close()

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
