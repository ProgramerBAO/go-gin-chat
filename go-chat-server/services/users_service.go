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
	"net/http"
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
// @param name formData string false "用户名"
// @param paw formData string false "密码"
// @param rePaw formData string false "核对密码"
// @param phoneNum formData string false "PhoneNum"
// @param email formData string false "Email"
// @Success 200 {string} Ok
// @Router /user/createUser [get]
func CreateUser(ctx *gin.Context) {
	var requestRegisterUser struct {
		UserName     string `json:"name"`
		UserPwd      string `json:"pwd"`
		UserRePwd    string `json:"rePwd"`
		UserPhoneNum string `json:"phoneNum"`
		UserEmail    string `json:"email"`
	}
	// 注意这里取地址
	// 直接判断
	fmt.Println(requestRegisterUser)
	// 将前端发送的JSON数据绑定到requestData结构体中
	if err := ctx.ShouldBindJSON(&requestRegisterUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
		return
	}
	user, err := models.FindUserByName(requestRegisterUser.UserName)
	user, err = models.FindUserByPhone(requestRegisterUser.UserPhoneNum)
	user, err = models.FindUserByEmail(requestRegisterUser.UserEmail)
	// 不存在记录
	if err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		ctx.JSON(-1, gin.H{
			"message": "姓名 电话 或者 邮箱已被占用",
		})
		return
	}

	user.Name = requestRegisterUser.UserName
	//passWorld := ctx.Query("paw")
	//rePassWorld := ctx.Query("rePaw")
	if requestRegisterUser.UserRePwd != requestRegisterUser.UserPwd {
		ctx.JSON(-1, gin.H{
			"message": "两次密码不一致,请检查",
		})
		return
	}
	fmt.Println("开始写入数据库")

	//获取一个随机数做盐值 很糙,需要更厉害的
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = utils.MakePassword(requestRegisterUser.UserPwd, salt)

	user.PhoneNum = requestRegisterUser.UserPhoneNum
	user.Email = requestRegisterUser.UserEmail
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

func GetUserByName(ctx *gin.Context) {
	var requestRegisterUser struct {
		// 这里json就是错的 帮不上
		UserName string `json:"name" form:"name"`
	}
	// 将前端发送的JSON数据绑定到requestData结构体中
	if err := ctx.ShouldBindQuery(&requestRegisterUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无效的请求数据",
		})
		return
	}
	fmt.Println("username=", ctx.Query("name"))
	fmt.Println("username2=", requestRegisterUser.UserName)

	user, err := models.FindUserByName(requestRegisterUser.UserName)
	fmt.Println("user= ", user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无效的请求数据",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"user": gin.H{
			"ID":   user.ID,
			"Name": user.Name,
		},
	})
}

func SendUserMsg(ctx *gin.Context) {
	Chat(ctx.Writer, ctx.Request)
}
