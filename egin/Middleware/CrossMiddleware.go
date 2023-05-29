package Middleware

import "github.com/gin-gonic/gin"

// CrossMiddleware 跨域中间件
func CrossMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域的Header
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")

		// 继续处理请求
		c.Next()
	}
}
