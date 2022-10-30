package impl

import (
	"context"
	"github.com/zxc-dev/micro_etcd_demo/login_service/model"
	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
)

type LoginService struct{}

var Db = model.DB

func (s *LoginService) Login(ctx context.Context, req *loginpb.LoginReq, rsp *loginpb.LoginRsp) error {
	var user model.User
	if Db.Where("name = ? and passwd = ?", req.Name, req.Passwd).First(&user).RecordNotFound() {
		rsp.Result = "用户名或密码错误，登录失败"
	} else {
		rsp.Result = "登录成功"
	}
	return nil
}
func (s *LoginService) Register(ctx context.Context, req *loginpb.RegisterReq, rsp *loginpb.RegisterRsp) error {
	if len(req.Name) > 45 {
		rsp.Result = "用户名过长，注册失败"
		return nil
	}
	if len(req.Passwd) > 45 {
		rsp.Result = "密码过长，注册失败"
		return nil
	}
	if len(req.Name) <= 2 || len(req.Passwd) <= 5 {
		rsp.Result = "注册失败，密码至少有5位，用户名至少为2个字符"
		return nil
	}

	var user model.User
	if Db.Where("name = ?", req.Name).First(&user).RecordNotFound() {
		user.Name = req.Name
		user.Passwd = req.Passwd
		Db.Create(&user)
		rsp.Result = "注册成功！"
	} else {
		rsp.Result = "注册失败，用户名已存在"
	}
	return nil
}
