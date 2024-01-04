package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-chat-server/models"
	"net/http"
)

// LoginService
// @Summary 登陆校验
// @Schemes
// @Description 登陆
// @Tags Users
// @Accept json
// @Produce application/json
// @param name query string false "用户名"
// @param pwd query string false "密码"
// @Success 200 {string} Ok
// @Router /login [post]
func LoginService(ctx *gin.Context) {
	// 密码要处理
	// 前端获取数据，用户名和密码，
	var requestUser struct {
		UserName string `json:"name"`
		UserPwd  string `json:"pwd"`
	}
	// 将前端发送的JSON数据绑定到requestData结构体中
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "无效的请求数据",
			"err":     err,
		})
		return
	}
	fmt.Println(requestUser)
	user, err := models.FindUserByNameAndPWD(requestUser.UserName, requestUser.UserPwd)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "账号或者密码错误",
		})
		return
	}
	// 加入鉴权
	models.AddJWT(user)
	//var friendList []models.Contact
	//if err := models.GetFriendList(&friendList, user.ID); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": "获取好友列表失败",
	//		"err":     err,
	//	})
	//	return
	//}
	ctx.JSON(200, gin.H{
		"message": "hello" + user.Name,
		"token":   user.Identity,
		"ID":      user.ID,
		//"friendList": friendList,
	})
}
