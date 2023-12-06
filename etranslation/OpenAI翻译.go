package etranslation

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/duolabmeng6/goefun/model/ejson"
	"io"
	"net/http"
	"strings"
)

// OpenAI翻译结构体
type OpenAI翻译 struct {
	Key      string
	Prompt   string
	E语言转转换键值 map[string]string
}

func NewOpenAI翻译(prompt, key string) *OpenAI翻译 {
	if prompt == "" {
		prompt = "Translate this into {目标语言}:\n\n{内容}\n\n"
	}
	return &OpenAI翻译{
		Prompt: prompt,
		Key:    key,
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
func (b *OpenAI翻译) E取初始化参数() []string {
	return []string{"prompt", "openai key"}
}

func (b *OpenAI翻译) E翻译(text, from, to string) (string, error) {
	if b.Key == "" {
		return "", errors.New("openid Key 为空")
	}

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	// 设置OpenAI API的访问密钥
	apiKey := b.Key
	client := &http.Client{}

	prompt := b.Prompt
	// 替换文本
	prompt = strings.Replace(prompt, "{目标语言}", to, -1)
	prompt = strings.Replace(prompt, "{内容}", text, -1)

	var data = make(map[string]any)
	data["model"] = "text-davinci-003"
	data["prompt"] = prompt
	data["temperature"] = 0.3
	data["max_tokens"] = 1000
	data["top_p"] = 1.0
	data["frequency_penalty"] = 0.0
	data["presence_penalty"] = 0.0
	jdata := new(bytes.Buffer)
	json.NewEncoder(jdata).Encode(data)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", jdata)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return "", err

	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Printf("%s\n", bodyText)
	翻译结果 := ejson.Json解析文本(string(bodyText), "choices.0.text")
	// 如果翻译结果 第一个是换行符 则删除第一个换行符
	if strings.HasPrefix(翻译结果, "\n") {
		翻译结果 = strings.TrimPrefix(翻译结果, "\n")
	}

	return 翻译结果, nil
}
