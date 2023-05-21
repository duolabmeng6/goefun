package egin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type header struct {
	Key   string
	Value string
}

func PerformRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {

	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func BindAndValid(c *gin.Context, params interface{}) error {
	var err error

	_ = c.ShouldBind(params) // 展示校验库，就先不多写err判断了
	// 校验
	//err := g.Validator().Data(params).Run(c)

	return err
}

func Test_jwttoken(t *testing.T) {
	r := gin.New()
	r.Any("/test", func(c *gin.Context) {
		type IndexRequest struct {
			Keywords string `i:"keywords" default:"" rule:"required|length:6,16" msg:"必填|名称长度为{min}到{max}个字符"`
			PerPage  int64  `i:"perPage" default:"" rule:"required" msg:"PerPage 必填"`
			Page     int64  `i:"page" default:"" rule:"required" msg:"Page 必填"`
			OrderBy  string `i:"orderby" default:""`
			OrderDir uint8  `i:"orderDir" default:""`
		}
		var uParam IndexRequest
		if err := Verify(c, &uParam); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"data": uParam,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": uParam,
		})
	})

	w := PerformRequest(r, http.MethodGet, "/test?page=1&perPage=10&keywords=123456")
	print(w.Body.String())

}

type IndexRequest struct {
	Keywords string `i:"keywords" default:"" rule:"required|length:6,16" msg:"必填|名称长度为{min}到{max}个字符"`
	PerPage  int64  `i:"perPage" default:"" rule:"required" msg:"PerPage 必填"`
	Page     int64  `i:"page" default:"" rule:"required" msg:"Page 必填"`
	OrderBy  string `i:"orderby" default:""`
	OrderDir uint8  `i:"orderDir" default:""`
}

func Index(c *gin.Context, req IndexRequest) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": req,
	})
}

func Test_控制器自动验证(t *testing.T) {
	r := gin.New()
	r.Any("/test", func(c *gin.Context) {
		var req IndexRequest
		if err := Verify(c, &req); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"data": req,
				"msg":  err.Error(),
			})
			return
		}
		Index(c, req)
	})

	w := PerformRequest(r, http.MethodGet, "/test?page=1&perPage=10&keywords=123456")
	print(w.Body.String())

}
