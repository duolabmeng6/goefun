package Middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/goefun/egin/jsonconv"
	"github.com/gin-gonic/gin"
	"strings"
)

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body      *bytes.Buffer // 缓存
	CamelCase bool          // 真驼峰 假下划线

}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	//如果是json格式
	if strings.Contains(w.Header().Get("Content-Type"), "application/json") {
		var u interface{}
		err := json.Unmarshal(b, &u)
		if err != nil {
			fmt.Sprintf("json.Unmarshal error %v", err)
			return 0, err
		}
		var marshal []byte
		if w.CamelCase {
			marshal, _ = json.Marshal(jsonconv.JsonCamelCase{Value: u})
		} else {
			marshal, _ = json.Marshal(jsonconv.JsonSnakeCase{Value: u})
		}

		b = []byte(fmt.Sprintf("%s", marshal))
	}

	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// JsonFormatMiddleware
//
// 自动将json数据重新转化为驼峰或者下划线
//
// 参数1 CamelCase true 驼峰 false 下划线
//
// 参数2 UrlPrefix 前缀 作用的url前缀
func JsonFormatMiddleware(CamelCase bool, UrlPrefix string) gin.HandlerFunc {
	// CamelCase true 驼峰 false 下划线
	return func(ctx *gin.Context) {
		//检查 UrlPrefix 前缀是否匹配 如果匹配则修改json返回格式
		print("ctx.Request.URL.Path=", ctx.Request.URL.Path, " UrlPrefix=", UrlPrefix, "\n")
		if strings.HasPrefix(ctx.Request.URL.Path, UrlPrefix) {
			w := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer, CamelCase: CamelCase}
			ctx.Writer = w
		}
		ctx.Next()
	}

}
