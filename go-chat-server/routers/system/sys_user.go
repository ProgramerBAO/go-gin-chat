package system

import (
	"github.com/gin-gonic/gin"
	"go-gin-chat-server/services"
	"go-gin-chat-server/utils"
)

type UserRouter struct {
}

func (receiver UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// userGroup := Router.Group("/user").Use(utils.JWTAuth())
	Router.Use(utils.VerifyToken())
	{
		Router.GET("/getUsers", services.GetUserByName) // 获得用户列表
		Router.GET("/deleteUser", services.DeleteUser)  // 删除用户
		Router.POST("/updateUser", services.UpdateUser) // 更新用户

		Router.GET("/addFriend", services.AddFriend)         // 添加好友
		Router.GET("/deleteFriend", services.DeleteFriend)   // 删除好友
		Router.GET("/getFriendList", services.GetFriendList) // 获取好友列表
	}
}
