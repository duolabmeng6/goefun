package etranslation

import (
	"errors"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/model/ejson"
)

// 彩云小译结构体
type 彩云小译 struct {
	Token    string
	E语言转转换键值 map[string]string
}

func New彩云小译(token string) *彩云小译 {
	return &彩云小译{
		Token: token,
		E语言转转换键值: map[string]string{
			"auto":                "auto",
			"Simplified Chinese":  "zh",
			"Traditional Chinese": "zh",
			"English":             "en",
			"Japanese":            "jp",
		},
	}
}
func (b *彩云小译) E取初始化参数() []string {
	return []string{"Token"}
}
func (b *彩云小译) E翻译(text, from, to string) (string, error) {
	if b.Token == "" {
		return "", errors.New("Token 为空")
	}

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	baseURL := "http://api.interpreter.caiyunai.com/v1/translator"
	postData := map[string]interface{}{
		"source":     text,
		"trans_type": from + "2" + to,
		"request_id": "demo",
		"detect":     true,
	}
	eh := ehttp.NewHttp()
	返回文本, err2 := eh.Post(baseURL, postData, map[string]string{
		"Content-Type":    "application/json",
		"x-authorization": "token " + b.Token,
	})
	//print(返回文本, err2)
	if err2 != nil {
		return "", errors.New("访问失败")
	}
	翻译结果 := ejson.Json解析文本(返回文本, "target")

	return 翻译结果, nil
}
