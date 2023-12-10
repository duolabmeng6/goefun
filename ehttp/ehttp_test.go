package ehttp

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func TestNewHttp(t *testing.T) {
	http := NewHttp()
	ret, _ := http.Get("https://www.taobao.com/help/getip.php")
	t.Log("访问失败", ret)
	//
	//retByte, flag := http.GetByte("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596964553726&di=e77c87e7f2c5a8b0b12bdf6c13fbefe9&imgtype=0&src=http%3A%2F%2Fa2.att.hudong.com%2F36%2F48%2F19300001357258133412489354717.jpg")
	//if flag {
	//	t.Log("访问失败", ret)
	//} else {
	//	t.Log(flag, retByte)
	//}
	//ecore.E写到文件("1.jpg", retByte)

}

func TestPOST(t *testing.T) {
	// 启动一个 gin
	go func() {
		r := gin.Default()
		r.Any("/test", func(c *gin.Context) {
			form := egin.IAll(c)
			header := c.Request.Header
			c.JSON(http.StatusOK, gin.H{
				"form":   form,
				"header": header,
			})
		})
		go r.Run(":8080")
	}()
	ecore.E延时(10)

	ehttp := NewHttp()
	ehttp.E设置全局HTTP代理("127.0.0.1:8888")
	//http.E设置全局HTTP代理("sockes5://192.168.100.1:7893")

	//ret, flag := http.Post("http://127.0.0.1:8080/test?ga=1&gb=2", "format=json&hasfast=true&authuser=0")
	//t.Log(flag, ret)
	//
	//ret, flag = http.Post("http://127.0.0.1:8080/test?ga=1&gb=2", `{"j1":1,"j2":"2"}`, "Content-Type: application/json")
	//t.Log(flag, ret)
	header := make(map[string]interface{})
	header["Content-Type"] = "application/json"
	header["abc"] = "abc"
	header["dddd"] = "abc"

	jsondata := map[string]interface{}{
		"j1": 1,
		"j2": 2,
		"j3": "abc",
	}

	ehttp.E设置全局头信息(header)
	//ret, flag := http.Post("http://127.0.0.1:8080/test?ga=1&gb=2", `{"j1":1,"j2":"2"}`, header)
	//t.Log(flag, ret)
	ret, flag := ehttp.Post("http://127.0.0.1:8080/test?ga=1&gb=2", jsondata)
	t.Log(flag, ret)
}

func testhttp() {
	tt := ecore.New时间统计类()

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.qq.com/", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(tt.E取毫秒())

}

func TestNewHttp2(t *testing.T) {
	ehttp := NewHttp()

	testhttp()
	testhttp()
	testhttp()
	testhttp()
	testhttp()

	ehttp.Get("https://www.qq.com/")
	ehttp.Get("https://www.qq.com/")
	ehttp.Get("https://www.qq.com/")
	ehttp.Get("https://www.qq.com/")
	ehttp.Get("https://www.qq.com/")

}
