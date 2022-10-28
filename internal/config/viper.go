package config

import (
	"github.com/spf13/viper"
)

// 全局Viper变量
var Viper = viper.New()

func Load(configPath string) error {
	Viper.SetConfigName("config")
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath(configPath)
	return Viper.ReadInConfig()
}
