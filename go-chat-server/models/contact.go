package models

import "gorm.io/gorm"

// Contact 这个表存放好友关系
type Contact struct {
	gorm.Model
	OwnerId  uint // 谁的关系信息 自己的微信号
	TargetId uint // 好友的微信号
	Type     int  // 对应的类型 好友 拉黑 用常量确定
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
