package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin-chat-server/docs"
	"go-gin-chat-server/routers/system"
	"go-gin-chat-server/services"
)

// Router 这里存放路由地址
func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//
	//// 静态资源
	//r.Static("asset", "asset/")
	//r.LoadHTMLGlob("views/**/*")

	// 首页
	r.GET("/", services.GetIndex)
	// 写到这里取写服务方法
	r.GET("/index", services.GetIndex)

	// 这个写法好
	sysRouterGroup := system.UseRouterGroup
	userGroup := r.Group("/user")
	{
		// 这个写法好
		sysRouterGroup.InitUserRouter(userGroup)
	}
	r.POST("/login", services.LoginService)
	// 注册用户
	userGroup.POST("/register", services.CreateUser)

	// 发送消息
	r.GET("/ws", services.SendMsg)
	r.GET("/sendUserMsg", services.SendUserMsg)

	// 测试主干配置修改工具
	r.POST("/new", services.ModifyConfig)

	return r
}
