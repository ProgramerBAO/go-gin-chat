package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"go-gin-chat-server/routers"
	"go-gin-chat-server/utils"
)

// 初始化函数
func init() {
	fmt.Println("开始运行函数,初始化数据生成中")
	// 初始化配置
	utils.InitConfig()
	utils.InitMySQL()
	fmt.Println("成功初始化数据库")
	utils.InitRedis()
	fmt.Println("成功初始化Redis")
}

func main() {
	r := routers.Router()
	r.Use(cors.Default())
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}
