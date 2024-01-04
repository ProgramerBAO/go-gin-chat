package models

import (
	"go-gin-chat-server/utils"
	"gorm.io/gorm"
)

// Contact 这个表存放好友关系
type Contact struct {
	gorm.Model
	OwnerId  uint   `json:"ownerId"`          // 谁的关系信息 自己的微信号
	TargetId uint   `json:"targetId"`         // 好友的微信号 能不呢存数组这样就可以节约空间
	Type     int    `json:"type" default:"1"` // 对应的类型 1 好友 2拉黑 用常量确定
	Desc     string `json:"desc"`             // 备注
}

func (table *Contact) TableName() string {
	return "contact"
}

// AddFriend 添加好友
func AddFriend(contact *Contact) error {
	return utils.DB.Create(contact).Error
}

// DeleteFriend 删除好友
func DeleteFriend(contact *Contact) error {
	return utils.DB.Where("owner_id = ? and target_id = ?", contact.OwnerId, contact.TargetId).Delete(&Contact{}).Error
}

// BlockFriend 拉黑好友
func BlockFriend(contact *Contact) *gorm.DB {
	return utils.DB.Model(contact).Update("type", 2)
}

// GetFriendList 获得好友列表
func GetFriendList(friendList *[]Contact, ownerId uint) error {
	res := utils.DB.Where("owner_id =?", ownerId).Find(&friendList)
	return res.Error
}
