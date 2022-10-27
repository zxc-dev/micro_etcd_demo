// greeter-api/client/greet.go

package client

import (
	"context"

	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
	"go-micro.dev/v4"
)

var loginClient loginpb.LoginService

// Init ... In the main function, you should call Init() first,
// so the 'client' defined above can be initialized.
func Init(s micro.Service) {
	// NewGreeterService is defined at proto/greeter.pb.micro.go file,
	// "micro.service.greeter" should match the name you defined in the server part.
	loginClient = loginpb.NewLoginService("micro.service.login", s.Client())
}

func Login(name, passwd string) (string, error) {
	res, err := loginClient.Login(context.TODO(), &loginpb.LoginRequest{
		Name:   name,
		Passwd: passwd,
	})

	if err != nil {
		return "", err
	}

	return res.Greeting, nil
}
