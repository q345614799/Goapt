package main

import (
	"apt/conf"
	"apt/init/initredis"
	"apt/init/initrouter"
	"apt/init/initsql"
	"fmt"
)


func main() {
	conf.Setviper()
	DB:= initsql.InitMysql(conf.Sqlconfig())
	defer DB.Close()
	initredis.InitClient()
	r:= initrouter.InitRouter()
	if err := r.Run(":9090");err!=nil{
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
