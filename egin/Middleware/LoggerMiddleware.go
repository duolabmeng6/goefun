package Middleware

import (
	"fmt"
	"github.com/duolabmeng6/goefun/elog"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	// 关闭gin的日志
	//gin.DisableConsoleColor()
	日志 := elog.New日志类("./log/gin.log", "info")

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		rawQuery := ""
		if c.Request.URL.RawQuery != "" {
			rawQuery = "?" + c.Request.URL.RawQuery
		}
		日志.E信息日志(
			"remoteIp", c.ClientIP(),
			"method", c.Request.Method,
			"url", c.Request.URL.Path+rawQuery,
			"status", c.Writer.Status(),
			"cost", fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		)
	}
}
