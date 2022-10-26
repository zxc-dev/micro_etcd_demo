package impl

import (
	"context"

	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
)

type LoginService struct{}

func (s *LoginService) Login(context.Context, *loginpb.LoginRequest, *loginpb.LoginResponse) error {
	panic("implement")

}
