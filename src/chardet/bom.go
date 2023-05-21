package chardet

// 使用bom来确定编码格式
func checkbom(data []byte) string {
	if len(data) >= 2 {
		if string(data[:2]) == "\xFE\xFF" {
			return "utf16-be"
		}
		if string(data[:2]) == "\xFF\xFE" {
			if len(data) >= 4 && string(data[2:4]) == "\x00\x00" {
				return "utf32-le"
			}
			return "utf16-le"
		}
	}
	if len(data) >= 3 {
		if string(data[:3]) == "\xEF\xBB\xBF" {
			return "utf8"
		}
	}
	if len(data) >= 4 {
		if string(data[:4]) == "\x00\x00\xFE\xFF" {
			return "utf32-be"
		}
		if string(data[:4]) == "\x84\x31\x95\x33" {
			return "gb18030"
		}
	}
	return ""
}
