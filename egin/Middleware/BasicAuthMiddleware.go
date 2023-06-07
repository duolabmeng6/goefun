package Middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// BasicAuthMiddleware http基本认证中间件
//
// Example:
//
//	r := gin.Default()
//	r.GET("/secret", BasicAuth("User Login", "admin", "admin"), func(c *gin.Context) {
//		c.String(200, "Secret content")
//	})
func BasicAuthMiddleware(realm string, username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取用户名和密码
		reqUsername, reqPassword, ok := c.Request.BasicAuth()
		if !ok {
			realm = strconv.Quote(realm)
			// 没有提供认证信息,返回401
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm="+realm)
			c.AbortWithStatus(401)

			return
		}

		// 验证用户名和密码
		if reqUsername != username || reqPassword != password {
			// 用户名或密码错误
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm="+realm)
			c.AbortWithStatus(401)
			return
		}

		// 认证成功,继续处理请求
		c.Next()
	}
}
