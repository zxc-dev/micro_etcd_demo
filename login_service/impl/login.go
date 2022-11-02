package impl

import (
	"context"
	loginpb "github.com/zxc-dev/micro_etcd_demo/pb/login"
)

type LoginService struct{}

type User struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"type:varchar(50);unique_index"`
	Passwd string `gorm:"type:varchar(50)"`
}

var users []User

/*var DB *gorm.DB

func InitMysql() {
	var err error
	DB, err = gorm.Open("mysql", "root:micro.demo.password@(mysql:3306)/micro_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err.Error())
	}

	DB.AutoMigrate(&User{})
	DB.Model(&User{}).AddUniqueIndex("name")
}
*/
func (s *LoginService) Login(ctx context.Context, req *loginpb.LoginReq, rsp *loginpb.LoginRsp) error {
	/*var user User
	if DB.Where("name = ? and passwd = ?", req.Name, req.Passwd).First(&user).RecordNotFound() {
		rsp.Result = "用户名或密码错误，登录失败"
	} else {
		rsp.Result = "登录成功"
	}*/
	for _, v := range users {
		if v.Name == req.Name && v.Passwd == req.Passwd {
			rsp.Result = "登陆成功！"
			return nil
		}
	}
	rsp.Result = "登录失败，请检查用户名和密码是否正确！"
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

	/*var user User
	if DB.Where("name = ?", req.Name).First(&user).RecordNotFound() {
		user.Name = req.Name
		user.Passwd = req.Passwd
		DB.Create(&user)
		rsp.Result = "注册成功！"
	} else {
		rsp.Result = "注册失败，用户名已存在"
	}*/
	user := User{
		Name:   req.Name,
		Passwd: req.Passwd,
	}
	for _, v := range users {
		if v.Name == user.Name {
			rsp.Result = "注册失败，用户名已存在！"
			return nil
		}
	}
	user.ID = int64(len(users))
	users = append(users, user)
	return nil
}
