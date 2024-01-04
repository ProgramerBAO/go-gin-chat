package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

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
	ind, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ind.Execute(ctx.Writer, "index")
	if err != nil {
		return
	}
	//ctx.JSON(200, gin.H{
	//	"meg": "welcome",
	//})
}
