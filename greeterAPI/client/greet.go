// greeter-api/client/greet.go

package client

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"greeterAPI/proto"
	"log"
)

var client proto.GreeterService

// Init ... In the main function, you should call Init() first,
// so the 'client' defined above can be initialized.
func Init() {
	//创建新服务
	service := micro.NewService(
		micro.Name("micro.client.greeter"),
	)
	//初始化
	service.Init()

	// NewGreeterService is defined at proto/greeter.pb.micro.go file,
	// "micro.service.greeter" should match the name you defined in the server part.
	client = proto.NewGreeterService("micro.service.greeter", service.Client())
}

// Greet ...
func Greet(ctx *gin.Context) {
	name := ctx.Query("name") // ctx.Query will return the GET request query string
	log.Println("Client handle Greet, name =", name)

	// Client request the RPC server for response
	res, err := client.Greet(context.TODO(), &proto.Request{Name: name})
	if err != nil {
		log.Print(err.Error())
		// return with 400 error code and error message
		ctx.JSON(400, gin.H{"message": "server error"})
		return
	}

	// return 200 success code and the response from server
	ctx.JSON(200, gin.H{"message": res.Greeting})
}