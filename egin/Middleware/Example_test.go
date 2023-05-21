package Middleware

import "github.com/gin-gonic/gin"

func ExampleJsonFormatMiddleware() {
	r := gin.Default()
	r.Use(JsonFormatMiddleware(false, "/admin"))
}
