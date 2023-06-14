package etranslation

import (
	"encoding/json"
	"strings"
	"sync"
)

type G语言列表 struct {
	E语言数组 []E语言列表结构
}

type E语言列表结构 struct {
	ChineseName      string `json:"chineseName"`
	EnglishName      string `json:"englishName"`
	LocalName        string `json:"localName"`
	FlagEmoji        string `json:"flagEmoji"`
	LocaleIdentifier string `json:"localeIdentifier"`
	VoiceName        string `json:"voiceName"`
}

var 语言列表json = `
[
  {
    "chineseName": "自动检测",
    "englishName": "Auto",
    "localName": "auto",
    "flagEmoji": "🌐",
    "localeIdentifier": "auto",
    "voiceName": "auto"
  },
  {
    "chineseName": "简体中文",
    "englishName": "Simplified Chinese",
    "localName": "简体中文",
    "flagEmoji": "🇨🇳",
    "localeIdentifier": "zh_CN",
    "voiceName": "Tingting"
  },
  {
    "chineseName": "繁体中文",
    "englishName": "Traditional Chinese",
    "localName": "繁體中文",
    "flagEmoji": "🇭🇰",
    "localeIdentifier": "zh_TW",
    "voiceName": "Tingting"
  },
  {
    "chineseName": "文言文",
    "englishName": "Classical Chinese",
    "localName": "文言文",
    "flagEmoji": "📜",
    "localeIdentifier": "zh_CN",
    "voiceName": "Tingting"
  },
  {
    "chineseName": "英语",
    "englishName": "English",
    "localName": "English",
    "flagEmoji": "🇬🇧",
    "localeIdentifier": "en_US",
    "voiceName": "Samantha"
  },
  {
    "chineseName": "日语",
    "englishName": "Japanese",
    "localName": "日本語",
    "flagEmoji": "🇯🇵",
    "localeIdentifier": "ja_JP",
    "voiceName": "Kyoko"
  },
  {
    "chineseName": "韩语",
    "englishName": "Korean",
    "localName": "한국어",
    "flagEmoji": "🇰🇷",
    "localeIdentifier": "ko_KR",
    "voiceName": "Yuna"
  },
  {
    "chineseName": "法语",
    "englishName": "French",
    "localName": "Français",
    "flagEmoji": "🇫🇷",
    "localeIdentifier": "fr_FR",
    "voiceName": "Amelie"
  },
  {
    "chineseName": "西班牙语",
    "englishName": "Spanish",
    "localName": "Español",
    "flagEmoji": "🇪🇸",
    "localeIdentifier": "es_ES",
    "voiceName": "Penelope"
  },
  {
    "chineseName": "葡萄牙语",
    "englishName": "Portuguese",
    "localName": "Português",
    "flagEmoji": "🇵🇹",
    "localeIdentifier": "pt_PT",
    "voiceName": "Luciana"
  },
  {
    "chineseName": "意大利语",
    "englishName": "Italian",
    "localName": "Italiano",
    "flagEmoji": "🇮🇹",
    "localeIdentifier": "it_IT",
    "voiceName": "Alice"
  },
  {
    "chineseName": "德语",
    "englishName": "German",
    "localName": "Deutsch",
    "flagEmoji": "🇩🇪",
    "localeIdentifier": "de_DE",
    "voiceName": "Anna"
  },
  {
    "chineseName": "俄语",
    "englishName": "Russian",
    "localName": "Русский",
    "flagEmoji": "🇷🇺",
    "localeIdentifier": "ru_RU",
    "voiceName": "Milena"
  },
  {
    "chineseName": "阿拉伯语",
    "englishName": "Arabic",
    "localName": "العربية",
    "flagEmoji": "🇸🇦",
    "localeIdentifier": "ar_AE",
    "voiceName": "Zuzana"
  },
  {
    "chineseName": "瑞典语",
    "englishName": "Swedish",
    "localName": "Svenska",
    "flagEmoji": "🇸🇪",
    "localeIdentifier": "sv_SE",
    "voiceName": "Alva"
  },
  {
    "chineseName": "罗马尼亚语",
    "englishName": "Romanian",
    "localName": "Română",
    "flagEmoji": "🇷🇴",
    "localeIdentifier": "ro_RO",
    "voiceName": "Ioana"
  },
  {
    "chineseName": "泰语",
    "englishName": "Thai",
    "localName": "ไทย",
    "flagEmoji": "🇹🇭",
    "localeIdentifier": "th_TH",
    "voiceName": "Kanya"
  },
  {
    "chineseName": "斯洛伐克语",
    "englishName": "Slovak",
    "localName": "Slovenčina",
    "flagEmoji": "🇸🇰",
    "localeIdentifier": "sk_SK",
    "voiceName": "Laura"
  },
  {
    "chineseName": "荷兰语",
    "englishName": "Dutch",
    "localName": "Nederlands",
    "flagEmoji": "🇳🇱",
    "localeIdentifier": "nl_NL",
    "voiceName": "Xander"
  },
  {
    "chineseName": "匈牙利语",
    "englishName": "Hungarian",
    "localName": "Magyar",
    "flagEmoji": "🇭🇺",
    "localeIdentifier": "hu_HU",
    "voiceName": "Ellen"
  },
  {
    "chineseName": "希腊语",
    "englishName": "Greek",
    "localName": "Ελληνικά",
    "flagEmoji": "🇬🇷",
    "localeIdentifier": "el_GR",
    "voiceName": "Melina"
  },
  {
    "chineseName": "丹麦语",
    "englishName": "Danish",
    "localName": "Dansk",
    "flagEmoji": "🇩🇰",
    "localeIdentifier": "da_DK",
    "voiceName": "Naja"
  },
  {
    "chineseName": "芬兰语",
    "englishName": "Finnish",
    "localName": "Suomi",
    "flagEmoji": "🇫🇮",
    "localeIdentifier": "fi_FI",
    "voiceName": "Satu"
  },
  {
    "chineseName": "波兰语",
    "englishName": "Polish",
    "localName": "Polski",
    "flagEmoji": "🇵🇱",
    "localeIdentifier": "pl_PL",
    "voiceName": "Ewa"
  },
  {
    "chineseName": "捷克语",
    "englishName": "Czech",
    "localName": "Čeština",
    "flagEmoji": "🇨🇿",
    "localeIdentifier": "cs_CZ",
    "voiceName": "Zuzana"
  },
  {
    "chineseName": "土耳其语",
    "englishName": "Turkish",
    "localName": "Türkçe",
    "flagEmoji": "🇹🇷",
    "localeIdentifier": "tr_TR",
    "voiceName": "Filiz"
  },
  {
    "chineseName": "立陶宛语",
    "englishName": "Lithuanian",
    "localName": "Lietuvių",
    "flagEmoji": "🇱🇹",
    "localeIdentifier": "lt_LT",
    "voiceName": "Rasa"
  },
  {
    "chineseName": "拉脱维亚语",
    "englishName": "Latvian",
    "localName": "Latviešu",
    "flagEmoji": "🇱🇻",
    "localeIdentifier": "lv_LV",
    "voiceName": "Liga"
  },
  {
    "chineseName": "乌克兰语",
    "englishName": "Ukrainian",
    "localName": "Українська",
    "flagEmoji": "🇺🇦",
    "localeIdentifier": "uk_UA",
    "voiceName": "Oksana"
  },
  {
    "chineseName": "保加利亚语",
    "englishName": "Bulgarian",
    "localName": "Български",
    "flagEmoji": "🇧🇬",
    "localeIdentifier": "bg_BG",
    "voiceName": "Tanya"
  },
  {
    "chineseName": "印尼语",
    "englishName": "Indonesian",
    "localName": "Bahasa Indonesia",
    "flagEmoji": "🇮🇩",
    "localeIdentifier": "id_ID",
    "voiceName": "Damayanti"
  },
  {
    "chineseName": "马来语",
    "englishName": "Malay",
    "localName": "Bahasa Melayu",
    "flagEmoji": "🇲🇾",
    "localeIdentifier": "ms_MY",
    "voiceName": "Zhiyu"
  },
  {
    "chineseName": "斯洛文尼亚语",
    "englishName": "Slovenian",
    "localName": "Slovenščina",
    "flagEmoji": "🇸🇮",
    "localeIdentifier": "sl_SI",
    "voiceName": "Lado"
  },
  {
    "chineseName": "爱沙尼亚语",
    "englishName": "Estonian",
    "localName": "Eesti",
    "flagEmoji": "🇪🇪",
    "localeIdentifier": "et_EE",
    "voiceName": "Karl"
  },
  {
    "chineseName": "越南语",
    "englishName": "Vietnamese",
    "localName": "Tiếng Việt",
    "flagEmoji": "🇻🇳",
    "localeIdentifier": "vi_VN",
    "voiceName": "An"
  },
  {
    "chineseName": "波斯语",
    "englishName": "Persian",
    "localName": "فارسی",
    "flagEmoji": "🇮🇷",
    "localeIdentifier": "fa_IR",
    "voiceName": "Zahra"
  },
  {
    "chineseName": "印地语",
    "englishName": "Hindi",
    "localName": "हिन्दी",
    "flagEmoji": "🇮🇳",
    "localeIdentifier": "hi_IN",
    "voiceName": "Lekha"
  },
  {
    "chineseName": "泰卢固语",
    "englishName": "Telugu",
    "localName": "తెలుగు",
    "flagEmoji": "🇮🇳",
    "localeIdentifier": "te_IN",
    "voiceName": "Chitra"
  },
  {
    "chineseName": "泰米尔语",
    "englishName": "Tamil",
    "localName": "தமிழ்",
    "flagEmoji": "🇮🇳",
    "localeIdentifier": "ta_IN",
    "voiceName": "Kanya"
  },
  {
    "chineseName": "乌尔都语",
    "englishName": "Urdu",
    "localName": "اردو",
    "flagEmoji": "🇮🇳",
    "localeIdentifier": "ur_PK",
    "voiceName": "Zaira"
  },
  {
    "chineseName": "菲律宾语",
    "englishName": "Filipino",
    "localName": "Filipino",
    "flagEmoji": "🇵🇭",
    "localeIdentifier": "fil_PH",
    "voiceName": ""
  },
  {
    "chineseName": "高棉语",
    "englishName": "Khmer",
    "localName": "ភាសាខ្មែរ",
    "flagEmoji": "🇰🇭",
    "localeIdentifier": "km_KH",
    "voiceName": ""
  },
  {
    "chineseName": "老挝语",
    "englishName": "Lao",
    "localName": "ພາສາລາວ",
    "flagEmoji": "🇱🇦",
    "localeIdentifier": "lo_LA",
    "voiceName": ""
  },
  {
    "chineseName": "孟加拉语",
    "englishName": "Bengali",
    "localName": "বাংলা",
    "flagEmoji": "🇧🇩",
    "localeIdentifier": "bn_BD",
    "voiceName": ""
  },
  {
    "chineseName": "缅甸语",
    "englishName": "Burmese",
    "localName": "ဗမာစာ",
    "flagEmoji": "🇲🇲",
    "localeIdentifier": "my_MM",
    "voiceName": ""
  },
  {
    "chineseName": "挪威语",
    "englishName": "Norwegian",
    "localName": "Norsk",
    "flagEmoji": "🇳🇴",
    "localeIdentifier": "nb_NO",
    "voiceName": ""
  },
  {
    "chineseName": "塞尔维亚语",
    "englishName": "Serbian",
    "localName": "Српски",
    "flagEmoji": "🇷🇸",
    "localeIdentifier": "sr_RS",
    "voiceName": ""
  },
  {
    "chineseName": "克罗地亚语",
    "englishName": "Croatian",
    "localName": "Hrvatski",
    "flagEmoji": "🇭🇷",
    "localeIdentifier": "hr_HR",
    "voiceName": ""
  },
  {
    "chineseName": "蒙古语",
    "englishName": "Mongolian",
    "localName": "Монгол",
    "flagEmoji": "🇲🇳",
    "localeIdentifier": "mn_MN",
    "voiceName": ""
  },
  {
    "chineseName": "希伯来语",
    "englishName": "Hebrew",
    "localName": "עברית",
    "flagEmoji": "🇮🇱",
    "localeIdentifier": "he_IL",
    "voiceName": ""
  }
]
`
var (
	instance *G语言列表
	once     sync.Once
)

func New语言列表() *G语言列表 {
	once.Do(func() {
		instance = &G语言列表{}
		json.Unmarshal([]byte(语言列表json), &instance.E语言数组)
	})
	return instance
}

func (c *G语言列表) E获取全部语言() []E语言列表结构 {
	return c.E语言数组
}

type E语言列表数据 struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (c *G语言列表) E取全部名称对照数组() []E语言列表数据 {
	// 生成数组 [{name: '自动检测', value: 'auto'}]

	arr := []E语言列表数据{}

	for _, 语言 := range c.E获取全部语言() {
		//println(语言.ChineseName, 语言.EnglishName, 语言.FlagEmoji)
		arr = append(arr, E语言列表数据{
			Name:  语言.FlagEmoji + 语言.ChineseName,
			Value: 语言.LocaleIdentifier,
		})
	}
	return arr
}

// E从区域标识取接口标识 从 zh_CN 匹配 EnglishName 到 转换结构体 的值
// 区域标识符 例如 zh_CN
// 转换结构体 例如
//
//	E语言转换接口: map[string]string{
//		"Auto":                "auto",
//		"Simplified Chinese":  "zh",
//		"Classical Chinese":   "wyw",
//		"Traditional Chinese": "cht",
//		"English":             "en",
//		"Japanese":            "jp",
func (c *G语言列表) E从区域标识取接口标识(区域标识符 string, 转换结构体 map[string]string) string {
	语言名称 := ""
	for _, 语言 := range c.E获取全部语言() {
		if 语言.LocaleIdentifier == 区域标识符 {
			语言名称 = 语言.EnglishName
			break
		}
	}
	if 语言名称 == "" {
		return 区域标识符
	}
	if 转换结构体 == nil {
		return 区域标识符
	}
	// 转换结构体 不要区分大小写获取key
	for k, v := range 转换结构体 {
		if strings.EqualFold(k, 语言名称) {
			return v
		}
	}

	return 区域标识符
}
