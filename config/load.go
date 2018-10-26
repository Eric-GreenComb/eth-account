package config

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/eth-account/bean"
)

// ServerConfig Server Config
var ServerConfig bean.ServerConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	ServerConfig.Port = strings.Split(viper.GetString("server.port"), ",")
	ServerConfig.Mode = viper.GetString("server.mode")
	ServerConfig.GormLogMode = viper.GetString("server.gorm.LogMode")
	ServerConfig.ViewLimit = viper.GetInt("server.view.limit")
}
