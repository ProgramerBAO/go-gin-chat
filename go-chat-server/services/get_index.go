package services

import "github.com/gin-gonic/gin"

// GetIndex godoc
// @Summary ping example
// @Schemes
// @Description 这是一个首页
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /index [get]
func GetIndex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"meg": "welcome",
	})
}
