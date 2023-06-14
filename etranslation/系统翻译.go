package etranslation

import (
	"fmt"
	"os/exec"
)

// 系统翻译结构体
type 系统翻译 struct {
	E语言转转换键值 map[string]string
}

func New系统翻译() *系统翻译 {
	return &系统翻译{
		E语言转转换键值: map[string]string{
			"auto":                "",
			"Simplified Chinese":  "zh_CN",
			"Traditional Chinese": "zh_TW",
			"English":             "en_US",
			"Japanese":            "ja_JP",
			"Korean":              "ko_KR",
			"French":              "fr_FR",
			"Spanish":             "es_ES",
			"Portuguese":          "pt_BR",
			"Italian":             "it_IT",
			"German":              "de_DE",
			"Russian":             "ru_RU",
			"Arabic":              "ar_AE",
			"Thai":                "th_TH",
			"Polish":              "pl_PL",
			"Turkish":             "tr_TR",
			"Indonesian":          "id_ID",
			"Vietnamese":          "vi_VN",
		},
	}
}
func (b *系统翻译) E取初始化参数() []string {
	return make([]string, 0)
}
func (b *系统翻译) E翻译(text, from, to string) (string, error) {
	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	// 构造AppleScript代码
	script := `
tell application "Shortcuts"
	set myInput to "text=` + text + "&from=" + from + `&to=` + to + `"
	set resultValue to run shortcut named "Easydict-Translate-V1.2.0" with input myInput
end tell
return resultValue
`
	//println(script)
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行AppleScript时出错:", err)
		return "", err
	}
	resultValue := string(output)
	return resultValue, nil
}
