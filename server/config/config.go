package config

import (
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func NewMysqlConfig() *MysqlConfig {

	//获取Mysql配置

	viper.SetConfigFile("/Users/hideoncode/Documents/RBAC/server/config/config.yaml")


	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &MysqlConfig{
		Host: viper.GetString("mysql.host"),
		Port: viper.GetString("mysql.port"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
	}
}