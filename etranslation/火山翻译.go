package etranslation

import (
	"errors"
	"github.com/duolabmeng6/goefun/etool"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

// 火山翻译结构体
type 火山翻译 struct {
	AppID    string
	Key      string
	E语言转转换键值 map[string]string
}

func New火山翻译(appID, key string) *火山翻译 {
	return &火山翻译{
		AppID: appID,
		Key:   key,
		E语言转转换键值: map[string]string{
			"Auto":                "auto",
			"Simplified Chinese":  "zh",
			"Traditional Chinese": "zh-Hant",
			"Classical Chinese":   "lzh",
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
			"Swedish":             "sv",
			"Romanian":            "ro",
			"Thai":                "th",
			"Slovak":              "sk",
			"Dutch":               "nl",
			"Hungarian":           "hu",
			"Greek":               "el",
			"Danish":              "da",
			"Finnish":             "fi",
			"Polish":              "pl",
			"Czech":               "cs",
			"Turkish":             "tr",
			"Lithuanian":          "lt",
			"Latvian":             "lv",
			"Ukrainian":           "uk",
			"Bulgarian":           "bg",
			"Indonesian":          "id",
			"Malay":               "ms",
			"Slovenian":           "sl",
			"Estonian":            "et",
			"Vietnamese":          "vi",
			"Persian":             "fa",
			"Hindi":               "hi",
			"Telugu":              "te",
			"Tamil":               "ta",
			"Urdu":                "ur",
			"Filipino":            "tl",
			"Khmer":               "km",
			"Lao":                 "lo",
			"Bengali":             "bn",
			"Burmese":             "my",
			"Norwegian":           "no",
			"Serbian":             "sr",
			"Croatian":            "hr",
			"Mongolian":           "mn",
			"Hebrew":              "iw",
		},
	}
}
func (b *火山翻译) E取初始化参数() []string {
	return []string{"AccessKey", "SecretKey"}
}
func (b *火山翻译) E翻译(text, from, to string) (string, error) {
	if b.AppID == "" || b.Key == "" {
		return "", errors.New("AppID 或者 Key 为空")
	}

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	kAccessKey := b.AppID
	kSecretKey := b.Key
	kServiceVersion := "2020-06-01"

	var (
		ServiceInfo = &base.ServiceInfo{
			Timeout: 5 * time.Second,
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "translate"},
		}
		ApiInfoList = map[string]*base.ApiInfo{
			"TranslateText": {
				Method: http.MethodPost,
				Path:   "/",
				Query: url.Values{
					"Action":  []string{"TranslateText"},
					"Version": []string{kServiceVersion},
				},
			},
		}
	)

	client := base.NewClient(ServiceInfo, ApiInfoList)
	client.SetAccessKey(kAccessKey)
	client.SetSecretKey(kSecretKey)

	//resp, code, err := client.Json("TranslateText", nil, `{"SourceLanguage":"`+from+`,"TargetLanguage":"`+to+`","TextList":["`+text+`"]}`)
	resp, _, err := client.Json("TranslateText", nil, `{"SourceLanguage":"`+from+`","TargetLanguage":"`+to+`","TextList":["`+text+`"]}`)
	if err != nil {
		return "", err
	}
	//fmt.Printf("%d %s\n", code, string(resp))

	翻译结果 := etool.Json解析文本(string(resp), "TranslationList.0.Translation")

	return 翻译结果, nil
}
