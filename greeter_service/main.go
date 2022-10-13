package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"greeter_service/proto"
	"log"
)

// GreeterService ...
type GreeterService struct {}

// Greet ... Implement interface left in proto/greeter.pb.micro.go server part
func (g *GreeterService) Greet(ctx context.Context, req *proto.Request, res *proto.Response) error {
	log.Println("Greeter service handle Greet", req.Name)
	res.Greeting = "Hello, " + req.Name
	return nil
}

func main() {
	reg := etcd.NewRegistry()

	service := micro.NewService(
		micro.Name("micro.service.greeter"), // The service name to register in the reg
		micro.Registry(reg),
	)

	service.Init()



	proto.RegisterGreeterHandler(service.Server(), &GreeterService{})

	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}