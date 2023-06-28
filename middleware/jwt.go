package middleware

import (
	"sdwebuiapi/e"
	"sdwebuiapi/utils/jwt"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTAuth token验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.AuthCheckTokenFailError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.AuthCheckTokenTimeoutError
			}
		}
		if code != e.SUCCESS {
			c.JSON(e.SUCCESS, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// JWTCheck 有token校验、无token不校验
func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token != "" {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.AuthCheckTokenFailError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.AuthCheckTokenTimeoutError
			}
		}
		if code != e.SUCCESS {
			c.JSON(e.SUCCESS, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
