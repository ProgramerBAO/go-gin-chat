package models

import "gorm.io/gorm"

// GroupBasic 这个表存放群信息
type GroupBasic struct {
	gorm.Model
	GroupName string // 群名称
	OwnerID   uint   // 群主ID
	Icon      string // 群头像
	Type      int    // 群类型
	Desc      string // 描述
	Members   []uint // 群成员ID 有问题

}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
