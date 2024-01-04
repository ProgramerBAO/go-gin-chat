package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-chat-server/models"
	"net/http"
	"strconv"
)

// AddFriend 添加好友
func AddFriend(ctx *gin.Context) {
	var contact = &models.Contact{
		OwnerId:  StrToUint(ctx.Query("ownerId")),
		TargetId: StrToUint(ctx.Query("targetId")),
		Type:     1,
		Desc:     ctx.Query("name"),
	}
	fmt.Println(contact)
	if contact.TargetId == 0 || contact.TargetId == contact.OwnerId {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "无效的请求数据",
			"err":     "targetId不能为空或您不能添加自己为好友",
		})
		return
	}
	if err := models.AddFriend(contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "添加好友失败",
			"err":     err,
		})
		return
	} else {
		var friendList []models.Contact
		if err := models.GetFriendList(&friendList, contact.OwnerId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "获取好友列表失败",
				"err":     err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "添加好友成功",
			"friendList": friendList,
		})
	}
}

// DeleteFriend 删除好友
func DeleteFriend(ctx *gin.Context) {
	var contact = &models.Contact{
		OwnerId:  StrToUint(ctx.Query("ownerId")),
		TargetId: StrToUint(ctx.Query("targetId")),
		Type:     1,
	}
	if err := models.DeleteFriend(contact); err != nil {
		fmt.Println("好友列表x", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "删除好友失败",
			"err":     err,
		})
		fmt.Println("好友列表x", err)
		return
	} else {
		var friendList []models.Contact
		fmt.Println("好友列表1", friendList)
		if err := models.GetFriendList(&friendList, contact.OwnerId); err != nil {
			fmt.Println("好友列表2", friendList)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "获取好友列表失败",
				"err":     err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "删除好友成功",
			"friendList": friendList,
		})
	}
}

// GetFriendList 获取好友列表
func GetFriendList(ctx *gin.Context) {
	var ownerId = StrToUint(ctx.Query("ownerId"))
	var friendList []models.Contact
	if err := models.GetFriendList(&friendList, ownerId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "获取好友列表失败",
			"err":     err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "获取好友列表成功",
		"friendList": friendList,
	})
}

// StrToUint string转换为uint
func StrToUint(str string) uint {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return uint(i)
}

// BlockFriend 拉黑好友
func BlockFriend(ctx *gin.Context) {

}
