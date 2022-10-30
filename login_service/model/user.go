package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"type:varchar(50);unique_index"`
	Passwd string `gorm:"type:varchar(50)"`
}

var DB *gorm.DB

func InitMysql() {
	var err error
	DB, err = gorm.Open("mysql", "root:micro.demo.password@(127.0.0.1:13306)/micro_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Errorf("创建数据库连接失败:%v", err)
	}

	DB.AutoMigrate(&User{})
	DB.Model(&User{}).AddUniqueIndex("uid", "name")
}
