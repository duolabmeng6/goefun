package etranslation

import (
	"errors"
	alimt20181012 "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/duolabmeng6/goefun/etool"
)

// 阿里云翻译结构体
type 阿里云翻译 struct {
	AppID    string
	Key      string
	E语言转转换键值 map[string]string
}

func New阿里云翻译(appID, key string) *阿里云翻译 {
	return &阿里云翻译{
		AppID: appID,
		Key:   key,
		E语言转转换键值: map[string]string{
			"Auto":                     "auto",
			"Simplified Chinese":       "zh",
			"Traditional Chinese":      "zh-tw",
			"English":                  "en",
			"Abkhazian":                "ab",
			"Albanian":                 "sq",
			"Akan":                     "ak",
			"Arabic":                   "ar",
			"Aragonese":                "an",
			"Amharic":                  "am",
			"Assamese":                 "as",
			"Azerbaijani":              "az",
			"Asturian":                 "ast",
			"Central Huasteca Nahuatl": "nch",
			"Ewe":                      "ee",
			"Aymara":                   "ay",
			"Irish":                    "ga",
			"Estonian":                 "et",
			"Ojibwa":                   "oj",
			"Occitan":                  "oc",
			"Oriya":                    "or",
			"Oromo":                    "om",
			"Ossetian":                 "os",
			"Tok Pisin":                "tpi",
			"Bashkir":                  "ba",
			"Basque":                   "eu",
			"Belarusian":               "be",
			"Berber languages":         "ber",
			"Bambara":                  "bm",
			"Pangasinan":               "pag",
			"Bulgarian":                "bg",
			"Northern Sami":            "se",
			"Bemba (Zambia)":           "bem",
			"Blin":                     "byn",
			"Bislama":                  "bi",
			"Baluchi":                  "bal",
			"Icelandic":                "is",
			"Polish":                   "pl",
			"Bosnian":                  "bs",
			"Persian":                  "fa",
			"Bhojpuri":                 "bho",
			"Breton":                   "br",
			"Chamorro":                 "ch",
			"Chavacano":                "cbk",
			"Chuvash":                  "cv",
			"Tsonga":                   "ts",
			"Tatar":                    "tt",
			"Danish":                   "da",
			"Shan":                     "shn",
			"Tetum":                    "tet",
			"German":                   "de",
			"Low German":               "nds",
			"Scots":                    "sco",
			"Dhivehi":                  "dv",
			"Kam":                      "kdx",
			"Kadazan Dusun":            "dtp",
			"Russian":                  "ru",
			"Faroese":                  "fo",
			"French":                   "fr",
			"Sanskrit":                 "sa",
			"Filipino":                 "fil",
			"Fijian":                   "fj",
			"Finnish":                  "fi",
			"Friulian":                 "fur",
			"Fur":                      "fvr",
			"Kongo":                    "kg",
			"Khmer":                    "km",
			"Guerrero Nahuatl":         "ngu",
			"Kalaallisut":              "kl",
			"Georgian":                 "ka",
			"Gronings":                 "gos",
			"Gujarati":                 "gu",
			"Guarani":                  "gn",
			"Kazakh":                   "kk",
			"Haitian":                  "ht",
			"Korean":                   "ko",
			"Hausa":                    "ha",
			"Dutch":                    "nl",
			"Montenegrin":              "cnr",
			"Hupa":                     "hup",
			"Gilbertese":               "gil",
			"Rundi":                    "rn",
			"K'iche'":                  "quc",
			"Kirghiz":                  "ky",
			"Galician":                 "gl",
			"Catalan":                  "ca",
			"Czech":                    "cs",
			"Kabyle":                   "kab",
			"Kannada":                  "kn",
			"Kanuri":                   "kr",
			"Kashubian":                "csb",
			"Khasi":                    "kha",
			"Cornish":                  "kw",
			"Xhosa":                    "xh",
			"Corsican":                 "co",
			"Creek":                    "mus",
			"Crimean Tatar":            "crh",
			"Klingon":                  "tlh",
			"Serbo-Croatian":           "hbs",
			"Quechua":                  "qu",
			"Kashmiri":                 "ks",
			"Kurdish":                  "ku",
			"Latin":                    "la",
			"Latgalian":                "ltg",
			"Latvian":                  "lv",
			"Lao":                      "lo",
			"Lithuanian":               "lt",
			"Limburgish":               "li",
			"Lingala":                  "ln",
			"Ganda":                    "lg",
			"Letzeburgesch":            "lb",
			"Rusyn":                    "rue",
			"Kinyarwanda":              "rw",
			"Romanian":                 "ro",
			"Romansh":                  "rm",
			"Romany":                   "rom",
			"Lojban":                   "jbo",
			"Malagasy":                 "mg",
			"Manx":                     "gv",
			"Maltese":                  "mt",
			"Marathi":                  "mr",
			"Malayalam":                "ml",
			"Malay":                    "ms",
			"Mari (Russia)":            "chm",
			"Macedonian":               "mk",
			"Marshallese":              "mh",
			"Kekchí":                   "kek",
			"Maithili":                 "mai",
			"Morisyen":                 "mfe",
			"Maori":                    "mi",
			"Mongolian":                "mn",
			"Bengali":                  "bn",
			"Burmese":                  "my",
			"Hmong":                    "hmn",
			"Umbundu":                  "umb",
			"Navajo":                   "nv",
			"Afrikaans":                "af",
			"Nepali":                   "ne",
			"Niuean":                   "niu",
			"Norwegian":                "no",
			"Pam":                      "pmn",
			"Papiamento":               "pap",
			"Panjabi":                  "pa",
			"Portuguese":               "pt",
			"Pushto":                   "ps",
			"Nyanja":                   "ny",
			"Twi":                      "tw",
			"Cherokee":                 "chr",
			"Japanese":                 "ja",
			"Swedish":                  "sv",
			"Samoan":                   "sm",
			"Sango":                    "sg",
			"Sinhala":                  "si",
			"Upper Sorbian":            "hsb",
			"Esperanto":                "eo",
			"Slovenian":                "sl",
			"Swahili":                  "sw",
			"Somali":                   "so",
			"Slovak":                   "sk",
			"Tagalog":                  "tl",
			"Tajik":                    "tg",
			"Tahitian":                 "ty",
			"Telugu":                   "te",
			"Tamil":                    "ta",
			"Thai":                     "th",
			"Tonga (Tonga Islands)":    "to",
			"Tonga (Zambia)":           "toi",
			"Tigrinya":                 "ti",
			"Tuvalu":                   "tvl",
			"Tuvinian":                 "tyv",
			"Turkish":                  "tr",
			"Turkmen":                  "tk",
			"Walloon":                  "wa",
			"Waray (Philippines)":      "war",
			"Welsh":                    "cy",
			"Venda":                    "ve",
			"Volapük":                  "vo",
			"Wolof":                    "wo",
			"Udmurt":                   "udm",
			"Urdu":                     "ur",
			"Uzbek":                    "uz",
			"Spanish":                  "es",
			"Interlingue":              "ie",
			"Western Frisian":          "fy",
			"Silesian":                 "szl",
			"Hebrew":                   "he",
			"Hiligaynon":               "hil",
			"Hawaiian":                 "haw",
			"Modern Greek":             "el",
			"Lingua Franca Nova":       "lfn",
			"Sindhi":                   "sd",
			"Hungarian":                "hu",
			"Shona":                    "sn",
			"Cebuano":                  "ceb",
			"Syriac":                   "syr",
			"Sundanese":                "su",
			"Armenian":                 "hy",
			"Achinese":                 "ace",
			"Iban":                     "iba",
			"Igbo":                     "ig",
			"Ido":                      "io",
			"Iloko":                    "ilo",
			"Inuktitut":                "iu",
			"Italian":                  "it",
			"Yiddish":                  "yi",
			"Interlingua":              "ia",
			"Hindi":                    "hi",
			"Indonesia":                "id",
			"Ingush":                   "inh",
			"Yoruba":                   "yo",
			"Vietnamese":               "vi",
			"Zaza":                     "zza",
			"Javanese":                 "jv",
			"Chinese":                  "zh",
			"Cantonese":                "yue",
			"Zulu":                     "zu",
		},
	}
}
func (b *阿里云翻译) E取初始化参数() []string {
	return []string{"Access Key ID", "Access Key Secret"}
}

func (b *阿里云翻译) E翻译(text, from, to string) (string, error) {
	if b.AppID == "" || b.Key == "" {
		return "", errors.New("AppID 或者 Key 为空")
	}

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	client, _err := createClient(tea.String(b.AppID), tea.String(b.Key))
	if _err != nil {
		return "", _err
	}

	translateGeneralRequest := &alimt20181012.TranslateGeneralRequest{
		TargetLanguage: tea.String(to),
		SourceLanguage: tea.String(from),
		SourceText:     tea.String(text),
		Scene:          tea.String("general"),
		FormatType:     tea.String("text"),
	}
	var ret *alimt20181012.TranslateGeneralResponse
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		ret, _err = client.TranslateGeneralWithOptions(translateGeneralRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return "", _err
		}
	}

	//println(ret.Body.String())
	翻译结果 := etool.Json解析文本(ret.Body.String(), "Data.Translated")

	return 翻译结果, nil
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *alimt20181012.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("mt.cn-hangzhou.aliyuncs.com")
	_result = &alimt20181012.Client{}
	_result, _err = alimt20181012.NewClient(config)
	return _result, _err
}
