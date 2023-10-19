package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("config")
	// 根据路径找到相应的配置项
	viper.SetConfigFile("config/app.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("app config", viper.Get("app"))
	fmt.Println("mysql config", viper.Get("mysql"))
}
