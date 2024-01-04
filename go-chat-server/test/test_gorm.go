package main

import (
	"fmt"
	"go-gin-chat-server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 测试数据库连接与增删改查功能
func main() {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/go_gin_chat_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("fail load db, error:%v\n", err)
		return
	}

	// 迁移schema
	// err = db.AutoMigrate(&models.UserBasic{})
	// err = db.AutoMigrate(&models.Message{})
	err = db.AutoMigrate(&models.GroupBasic{})
	//err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		fmt.Println(err)
		return
	}

	//userIn := &models.UserBasic{
	//	Model: gorm.Model{},
	//	Name:  "BobShen",
	//}
	//// 增
	//db.Create(userIn)
	//// 查(取)
	//var userOut models.UserBasic
	//db.First(&userOut, 1)
	//fmt.Println(userOut)
	//
	//// 改
	//db.Model(userIn).Update("Name", "Kenny")
	//db.First(&userOut, 1)
	//fmt.Println(userOut)
	//// 删
	//// db.Delete(&userIn,1)
}
