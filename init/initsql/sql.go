package initsql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DEFAULTDB *gorm.DB

//初始化数据库并产生数据库全局变量
func InitMysql(config string) *gorm.DB {

	db,err:= gorm.Open("mysql", config)
	if err!=nil {
		fmt.Printf("数据库连接错误：%#V",err.Error())
	}else {
		DEFAULTDB=db
		fmt.Printf("数据库连接成功")
	}

	return DEFAULTDB

}
