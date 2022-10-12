package main

import (
	"context"
	"go-micro.dev/v4"
	"greeterService/proto"
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
	service := micro.NewService(
		micro.Name("micro.service.greeter"), // The service name to register in the registry
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), &GreeterService{})

	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}