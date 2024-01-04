package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// 这里存放jwt的相关的
type myClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username, pwd string) string {
	// 我的理解是类似盐值 这里就用用户名吧 一旦用户名密码修改就失效
	mySigningKey := []byte("username")
	c := myClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发人
			Issuer:   username,
			Subject:  "",
			Audience: nil,
			// 过期时间 当前时间的一分钟之后
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			// 开始生效 当前时间的一分钟之前
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ID:       "",
		},
	}
	// 生成了token 未加密的版本
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println("token=", token)
	// 这个才是最终的token 加密后的版本
	signedString, err := token.SignedString(mySigningKey)
	fmt.Println("signedString=", signedString)
	if err != nil {
		fmt.Println("生成token失败", err)
		return ""
	}
	// fmt.Println(signedString)
	// 解密
	token2, err := jwt.ParseWithClaims(signedString, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("token=", token)
		fmt.Println("token2=", token2)
	}
	return signedString
}

// VerifyToken 验证token的有效性
func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取现有的token
		token := ctx.Request.Header.Get("x-token")
		// 验证token是否有效
		// 1 没有token
		if token == "" {
			fmt.Println("没有token")
			// 这里要返回
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的token",
			})
			ctx.Abort()
			return
		}

		fmt.Println("取到token了", token)
		// 解密 token  模板 还有之前的的key  我们使用的是username和密码
		token2, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("username"), nil
		})
		// 无效的处理方式
		if err != nil {
			fmt.Println("出现了未知错误", err)
			// 无效token
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的token",
			})
			ctx.Abort()
			return
		}
		// 有效的处理方式
		fmt.Println(token2.Claims.(*myClaims).Username)
		ctx.Next()
	}
}
