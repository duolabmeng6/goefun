package etranslation

import "regexp"

func E内容语言类型检测(text string) string {
	// 匹配日文
	if regexp.MustCompile(`[\p{Hiragana}\p{Katakana}]`).MatchString(text) {
		return "ja_JP"
	}
	// 匹配中文
	if regexp.MustCompile(`[\p{Han}]`).MatchString(text) {
		return "zh_CN"
	}
	// 匹配韩文
	if regexp.MustCompile(`[\p{Hangul}]`).MatchString(text) {
		return "ko_KR"
	}

	// 匹配英文和其他语言
	return "en_US"
}

func E简写转中文(abbreviation string) string {
	switch abbreviation {
	case "zh_CN":
		return "中文"
	case "en_US":
		return "英文"
	case "ja_JP":
		return "日文"
	case "ko_KR":
		return "韩文"
	default:
		return "英文"
	}
}
