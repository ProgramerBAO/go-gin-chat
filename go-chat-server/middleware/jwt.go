package middleware

import "github.com/gin-gonic/gin"

// 这个文件是为了校验jwt 登陆页面的时候会生成并重置jwt

func VerifiedJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Header.Get("x-token")
	}
}
