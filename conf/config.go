package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type sqlconfig struct {
	Addr string
	Port string
	Username string
	Password string
	Database string
}
type redisconfig struct {
	Addr string
	Port string
}


func Sqlconfig() string {
	var config sqlconfig
	config.Addr =viper.GetString("Addr")
	config.Port =viper.GetString("Port")
	config.Username =viper.GetString("Username")
	config.Password =viper.GetString("Password")
	config.Database =viper.GetString("Database")
	var mysqlconfig string
	mysqlconfig = config.Username+":"+config.Password+"@("+config.Addr+":"+config.Port+")/"+config.Database+"?charset=utf8mb4&parseTime=True&loc=Local"
	return mysqlconfig
}

func Redisconfig() string {
	var config redisconfig
	config.Addr =viper.GetString("Addr")
	config.Port =viper.GetString("Port")
	var redisconfig string
	redisconfig = config.Addr+":"+config.Port
	return redisconfig
}

func Setviper()  {
	viper.SetConfigName("config")  // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")    // 指定配置文件类型
	viper.AddConfigPath("./conf/") // 指定查找配置文件的路径（这里使用相对路径）
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("未找到配置文件: %s \n", err))
		} else {
			panic(fmt.Errorf("发生意外错误: %s \n", err))
		}
	}
}
