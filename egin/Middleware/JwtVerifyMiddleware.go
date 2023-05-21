package Middleware

import (
	"github.com/duolabmeng6/goefun/egin/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// JwtVerifyMiddleware
// @description: jwt验证中间件 未授权跳转到登录页面 已授权将uid写入上下文 通过c.Get("uid")获取
// @param RedirectUrl string 未授权跳转地址 例如 /admin/login
// @return gin.HandlerFunc
func JwtVerifyMiddleware(RedirectUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			//从 cookie 中获取  jwt
			authorization, _ = c.Cookie("jwt")
		}
		if authorization == "" {
			//跳转到登录页面
			c.Redirect(302, RedirectUrl)
			c.Abort()
			return
		}
		auth, valid := jwt.ParseToken(authorization)
		if valid {
			c.Set("uid", strconv.FormatInt(auth.Uid, 10))
			c.Next()
		} else {
			c.Redirect(302, RedirectUrl)
			c.Abort()
		}
	}
}
