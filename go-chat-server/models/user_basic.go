package models

import (
	"database/sql"
	"fmt"
	"go-gin-chat-server/utils"
	"gorm.io/gorm"
)

// 数据库

type UserBasic struct {
	gorm.Model
	Name          string       `gorm:"column:name"`
	Password      string       `gorm:"column:password"`
	PhoneNum      string       `gorm:"column:phone_num" valid:"matches(^1[3456789]\\d{9}$)"`
	Email         string       `gorm:"column:email" valid:"email"`
	Identity      string       `gorm:"column:identity"`
	ClientIP      string       `gorm:"column:client_ip"`
	ClientPort    string       `gorm:"column:client_port"`
	LoginTime     sql.NullTime `gorm:"column:login_time"`
	HeartBeatTime sql.NullTime `gorm:"column:heart_beat_time"`
	LogoutTime    sql.NullTime `gorm:"column:logout_time"`
	IsLogout      bool         `gorm:"column:is_logout"`
	Device        string       `gorm:"column:device"`
	Salt          string       `gorm:"column:salt"`
}

// TableName 实体类  user_basic是表名
func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	userList := make([]*UserBasic, 10)
	utils.DB.Find(&userList)

	for _, basic := range userList {
		fmt.Println(basic)
	}
	return userList
}

func CreateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Create(user)
}

func DeleteUserByName(user *UserBasic) *gorm.DB {
	return utils.DB.Where("name = ?", user.Name).Delete(user)
}

func UpdateUser(user *UserBasic) {
	// return utils.DB.Model(user).Where("id = ?", user.ID).Update("name", user.Name)
	// 更新多个字段
	utils.DB.Model(user).Updates(user)
}

// AddJWT 将token增加到用户中,
func AddJWT(basic *UserBasic) {
	basic.Identity = utils.GenerateJWT(basic.Name, basic.Password)
	fmt.Println("basic.Identity:", basic.Identity)
	UpdateUser(basic)
}

// FindUserByName 查找
func FindUserByName(name string) (*UserBasic, error) {
	// 名称应该是唯一的
	user := &UserBasic{}
	result := utils.DB.Where("name = ?", name).First(user)
	return user, result.Error
}

// FindUserByPhone 根据手机号查找
func FindUserByPhone(phoneNum string) (*UserBasic, error) {
	// 手机号应该是唯一的
	user := &UserBasic{}
	result := utils.DB.Where("phone_num = ?", phoneNum).First(user)
	return user, result.Error
}

// FindUserByEmail 根据邮箱查找
func FindUserByEmail(email string) (*UserBasic, error) {
	// 邮箱应该是唯一的
	user := &UserBasic{}
	result := utils.DB.Where("email = ?", email).First(user)
	return user, result.Error
}
