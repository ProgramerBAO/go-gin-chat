package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ModifyConfig(ctx *gin.Context) {
	var requestData struct {
		NewInfo      string `json:"newInfo"`
		UpdateInfo   string `json:"updateInfo"`
		OriginalInfo string `json:"originalInfo"`
	}

	// 将前端发送的JSON数据绑定到requestData结构体中
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
		return
	}

	// 在这里可以使用 requestData 中的三个信息执行后端逻辑
	// 示例中仅返回一个响应消息
	responseMessage := fmt.Sprintf("成功接收数据: 新增信息=%s, 修改信息=%s, 原始信息=%s", requestData.NewInfo, requestData.UpdateInfo, requestData.OriginalInfo)
	ctx.JSON(http.StatusOK, gin.H{"message": responseMessage})
}
