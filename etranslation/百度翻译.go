package etranslation

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/etool"
	"net/url"
	"strconv"
)

// 百度翻译结构体
type 百度翻译 struct {
	AppID    string
	Key      string
	E语言转转换键值 map[string]string
}

func New百度翻译(appID, key string) *百度翻译 {
	return &百度翻译{
		AppID: appID,
		Key:   key,
		E语言转转换键值: map[string]string{
			"Auto":                "auto",
			"Simplified Chinese":  "zh",
			"Classical Chinese":   "wyw",
			"Traditional Chinese": "cht",
			"English":             "en",
			"Japanese":            "jp",
			"Korean":              "kor",
			"French":              "fra",
			"Spanish":             "spa",
			"Portuguese":          "pt",
			"Italian":             "it",
			"German":              "de",
			"Russian":             "ru",
			"Arabic":              "ara",
			"Swedish":             "swe",
			"Romanian":            "rom",
			"Thai":                "th",
			"Slovak":              "slo",
			"Dutch":               "nl",
			"Hungarian":           "hu",
			"Greek":               "el",
			"Danish":              "dan",
			"Finnish":             "fin",
			"Polish":              "pl",
			"Czech":               "cs",
			"Turkish":             "tr",
			"Lithuanian":          "lit",
			"Latvian":             "lav",
			"Ukrainian":           "ukr",
			"Bulgarian":           "bul",
			"Indonesian":          "id",
			"Malay":               "msa",
			"Slovenian":           "slv",
			"Estonian":            "est",
			"Vietnamese":          "vie",
			"Persian":             "per",
			"Hindi":               "hin",
			"Telugu":              "tel",
			"Tamil":               "tam",
			"Urdu":                "urd",
			"Filipino":            "fil",
			"Khmer":               "khm",
			"Lao":                 "lo",
			"Bengali":             "ben",
			"Burmese":             "bur",
			"Norwegian":           "nor",
			"Serbian":             "srp",
			"Croatian":            "hrv",
			"Mongolian":           "mon",
			"Hebrew":              "heb",
		},
	}
}
func (b *百度翻译) E取初始化参数() []string {
	return []string{"appid", "secret"}
}
func (b *百度翻译) E翻译(text, from, to string) (string, error) {
	if b.AppID == "" || b.Key == "" {
		return "", errors.New("AppID 或者 Key 为空")
	}

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)
	//print("form", from, "to", to, "\n")
	// 设置 API 请求参数
	q := text
	fromLang := from
	toLang := to
	salt := ecore.E取随机数(32768, 65536)
	sign := b.AppID + q + strconv.Itoa(salt) + b.Key
	hash := md5.Sum([]byte(sign))
	sign = hex.EncodeToString(hash[:])

	baseURL := "http://api.fanyi.baidu.com/api/trans/vip/translate"
	url, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	params := url.Query()
	params.Set("appid", b.AppID)
	params.Set("q", q)
	params.Set("from", fromLang)
	params.Set("to", toLang)
	params.Set("salt", strconv.Itoa(salt))
	params.Set("sign", sign)
	url.RawQuery = params.Encode()

	eh := ehttp.NewHttp()
	返回文本, err2 := eh.Get(url.String())
	if err2 != nil {
		return "", errors.New("访问失败")
	}
	翻译结果 := etool.Json解析文本(返回文本, "trans_result.0.dst")

	return 翻译结果, nil
}
