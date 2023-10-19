package services

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go-gin-chat-server/models"
	"go-gin-chat-server/utils"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

// GetUsers godoc
// @Summary 获取用户列表
// @Schemes
// @Description 用户模块
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} HelloWorld
// @Router /user/getUsers [get]
func GetUsers(ctx *gin.Context) {
	user := &models.UserBasic{}
	utils.DB.First(user)
	userJson, _ := json.Marshal(user)
	ctx.JSON(200, gin.H{
		"user": string(userJson),
	})
}

// CreateUser godoc
// @Summary 新建用户
// @Schemes
// @Description 用户模块
// @Tags Users
// @Accept json
// @Produce json
// @param name query string false "用户名"
// @param paw query string false "密码"
// @param rePaw query string false "核对密码"
// @param phoneNum query string false "PhoneNum"
// @param email query string false "Email"
// @Success 200 {string} Ok
// @Router /user/createUser [get]
func CreateUser(ctx *gin.Context) {
	// 注意这里取地址
	// 直接判断
	user, err := models.FindUserByName(ctx.Query("name"))
	user, err = models.FindUserByPhone(ctx.Query("phoneNum"))
	user, err = models.FindUserByEmail(ctx.Query("email"))
	// 不存在记录
	if err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		ctx.JSON(-1, gin.H{
			"message": "姓名 电话 或者 邮箱已被占用",
		})
		return
	}

	user.Name = ctx.Query("name")
	passWorld := ctx.Query("paw")
	rePassWorld := ctx.Query("rePaw")
	if passWorld != rePassWorld {
		ctx.JSON(-1, gin.H{
			"message": "两次密码不一致,请检查",
		})
		return
	}
	fmt.Println("开始写入数据库")

	//获取一个随机数做盐值 很糙,需要更厉害的
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = utils.MakePassword(passWorld, salt)

	user.PhoneNum = ctx.Query("phoneNum")
	user.Email = ctx.Query("email")
	models.CreateUser(user)
	ctx.JSON(200, gin.H{
		"message": "添加成功",
	})
}

// DeleteUser godoc
// @Summary 删除用户(逻辑删除)
// @Schemes
// @Description 用户模块
// @Tags Users
// @param name query string false "用户名"
// @Success 200 {string} Ok
// @Router /user/deleteUser [get]
func DeleteUser(ctx *gin.Context) {
	// 注意这里取地址
	user := &models.UserBasic{}
	user.Name = ctx.Query("name")
	// 根据用户名 但是不安全
	// utils.DB.Where("name = ?", ctx.Query("name")).Delete(user)
	models.DeleteUserByName(user)
	ctx.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags Users
// @param name formData string false "用户名"
// @param id formData string false "ID"
// @param phoneNum formData string false "PhoneNum"
// @param email formData string false "Email"
// @Success 200 {string} Ok
// @Router /user/updateUser [post]
func UpdateUser(ctx *gin.Context) {
	// 注意这里取地址
	user := &models.UserBasic{}
	user.Name = ctx.PostForm("name")
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		fmt.Println("转换失败 ", err)
		return
	}
	user.ID = uint(id)
	user.PhoneNum = ctx.PostForm("phoneNum")
	user.Email = ctx.PostForm("email")
	// 后端校验
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, gin.H{
			"message": "更新失败",
		})
		return
	}
	// 根据用户名 但是不安全
	fmt.Println("开始写入数据库")
	models.UpdateUser(user)
	ctx.JSON(200, gin.H{
		"message": "更新成功",
	})
}

func SendUserMsg(ctx *gin.Context) {
	Chat(ctx.Writer, ctx.Request)
}
