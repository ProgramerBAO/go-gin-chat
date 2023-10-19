package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin-chat-server/docs"
	"go-gin-chat-server/services"
)

// Router 这里存放路由地址
func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 写到这里取写服务方法
	r.GET("/index", services.GetIndex)
	r.GET("/user/getUsers", services.GetUsers)
	r.GET("/user/createUser", services.CreateUser)
	r.GET("/user/deleteUser", services.DeleteUser)
	r.POST("/user/updateUser", services.UpdateUser)
	r.POST("/login", services.LoginService)

	// 发送消息
	r.GET("/ws", services.SendMsg)
	r.GET("/sendUserMsg", services.SendUserMsg)

	// 测试主干配置修改工具
	r.POST("/new", services.ModifyConfig)

	return r
}
