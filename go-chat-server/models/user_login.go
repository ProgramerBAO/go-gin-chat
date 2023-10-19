package models

import (
	"go-gin-chat-server/utils"
	"gorm.io/gorm"
)

func FindUserByNameAndPWD(name, pwd string) (*UserBasic, error) {
	// 登录 查询用户名和密码是否匹配
	user := &UserBasic{}
	result := utils.DB.Where("name = ?", name).First(user)
	// 记录存在
	if result.Error != gorm.ErrRecordNotFound {
		realPwd := utils.MakePassword(pwd, user.Salt)
		result = utils.DB.Where("name = ? and password = ?", name, realPwd).First(user)
	}
	return user, result.Error
}
