package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	// 设置 Viper 解析 YAML 配置文件
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

}
