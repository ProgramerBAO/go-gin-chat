package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type myClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username, pwd string) string {
	// 我的理解是累死盐值 这里就用用户名吧 一旦用户名密码修改就失效
	mySigningKey := []byte(username + pwd)
	c := myClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发人
			Issuer:   username,
			Subject:  "",
			Audience: nil,
			// 过期时间 当前时间的一分钟之后
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
			// 开始生效 当前时间的一分钟之前
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ID:       "",
		},
	}
	// 生成了token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(mySigningKey)
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
		fmt.Println("token1", token)
		fmt.Println("token2", token2)
	}
	return signedString
}
