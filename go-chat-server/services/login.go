package services

import (
	"github.com/gin-gonic/gin"
	"go-gin-chat-server/models"
)

// LoginService
// @Summary 登陆校验
// @Schemes
// @Description 登陆
// @Tags Users
// @Accept json
// @Produce json
// @param name query string false "用户名"
// @param pwd query string false "密码"
// @Success 200 {string} Ok
// @Router /login [post]
func LoginService(ctx *gin.Context) {
	// 密码要处理
	// 前端获取数据，用户名和密码，
	user, err := models.FindUserByNameAndPWD(ctx.Query("name"), ctx.Query("pwd"))
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "账号或者密码错误",
		})
		return
	}
	// 加入鉴权
	models.AddJWT(user)
	ctx.JSON(200, gin.H{
		"message": "hello" + user.Name,
		"token":   user.Identity,
	})
}
