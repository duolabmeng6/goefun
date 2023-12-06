package etranslation

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/model/ejson"
	"github.com/google/uuid"
	"net/url"
	"time"
)

// 有道翻译结构体
type 有道翻译 struct {
	AppID    string
	Key      string
	E语言转转换键值 map[string]string
}

func New有道翻译(appID, key string) *有道翻译 {
	return &有道翻译{
		AppID: appID,
		Key:   key,
		E语言转转换键值: map[string]string{
			"Simplified Chinese":  "zh-CHS",
			"Traditional Chinese": "zh-CHT",
			"English":             "en",
			"Japanese":            "ja",
			"Korean":              "ko",
			"French":              "fr",
			"Spanish":             "es",
			"Portuguese":          "pt",
			"Italian":             "it",
			"German":              "de",
			"Russian":             "ru",
			"Arabic":              "ar",
			"Thai":                "th",
			"Dutch":               "nl",
			"Indonesian":          "id",
			"Vietnamese":          "vi",
		},
	}
}
func (b *有道翻译) E取初始化参数() []string {
	return []string{"appKey", "appSecret"}
}
func (y *有道翻译) Encrypt(signStr string) string {
	hashAlgorithm := sha256.New()
	hashAlgorithm.Write([]byte(signStr))
	return hex.EncodeToString(hashAlgorithm.Sum(nil))
}

func (y *有道翻译) Truncate(q string) string {
	size := len(q)
	if size <= 20 {
		return q
	}
	return q[:10] + fmt.Sprintf("%d", size) + q[size-10:size]
}

func (b *有道翻译) E翻译(text, from, to string) (string, error) {
	if b.AppID == "" || b.Key == "" {
		return "", errors.New("AppID 或者 Key 为空")
	}

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	data := url.Values{}
	data.Set("from", from)
	data.Set("to", to)
	data.Set("signType", "v3")
	curtime := fmt.Sprintf("%d", time.Now().Unix())
	data.Set("curtime", curtime)
	salt := uuid.New().String()
	signStr := b.AppID + b.Truncate(text) + salt + curtime + b.Key
	sign := b.Encrypt(signStr)
	data.Set("appKey", b.AppID)
	data.Set("q", text)
	data.Set("salt", salt)
	data.Set("sign", sign)

	eh := ehttp.NewHttp()
	返回文本, err2 := eh.Post("https://openapi.youdao.com/api", data.Encode())
	if err2 != nil {
		return "", errors.New("访问失败")
	}
	print(返回文本)
	翻译结果 := ejson.Json解析文本(返回文本, "translation")

	return 翻译结果, nil
}
