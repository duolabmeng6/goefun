// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供编码转换功能
package ecore

import (
	"github.com/axgle/mahonia"
	"github.com/duolabmeng6/goefun/src/chardet"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gurl"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// Base64编码 将数据编码为Base64格式。
//
// 参数：
//   - data: 要编码的字节集数据
//
// 返回值：
//   - string: Base64编码后的字符串
//
// 示例：
//
//	Base64编码([]byte("hello")) // 返回 "aGVsbG8="
func Base64编码(data []byte) string {
	return gbase64.EncodeToString(data)
}

// Base64解码 将Base64文本解码为原始数据。
//
// 参数：
//   - data: 要解码的Base64文本
//
// 返回值：
//   - string: 解码后的原始文本
//
// 示例：
//
//	Base64解码("aGVsbG8=") // 返回 "hello"
func Base64解码(data string) string {
	str, _ := gbase64.DecodeToString(data)
	return str
}

// URL编码 对URL内容进行编码。
//
// 参数：
//   - str: 要进行URL编码的文本内容
//
// 返回值：
//   - string: 编码后的URL字符串
//
// 示例：
//
//	URL编码("你好世界") // 返回 "%E4%BD%A0%E5%A5%BD%E4%B8%96%E7%95%8C"
func URL编码(str string) string {
	return gurl.Encode(str)
}

// URL解码 对URL内容进行解码。
//
// 参数：
//   - str: 要进行URL解码的文本内容
//
// 返回值：
//   - string: 解码后的原始文本
//
// 示例：
//
//	URL解码("%E4%BD%A0%E5%A5%BD") // 返回 "你好"
func URL解码(str string) string {
	s, _ := gurl.Decode(str)
	return s
}

// URL解析 解析URL并返回其组成部分。
//
// 参数：
//   - str: 要解析的URL字符串
//   - component: 解析组成部分：-1表示全部；1表示scheme；2表示host；4表示port；8表示user；16表示pass；32表示path；64表示query；128表示fragment
//
// 返回值：
//   - map[string]string: URL各组成部分的映射
//
// 示例：
//
//	URL解析("http://user:pass@example.com:8080/path", -1)
func URL解析(str string, component int) map[string]string {
	s, _ := gurl.ParseURL(str, component)
	return s
}

// E文本编码转换 将文本从一种编码转换为另一种编码。
//
// 参数：
//   - str: 要转换的文本
//   - 来源编码: 原始编码（如"gbk"、"utf-8"），为空则自动检测
//   - 目标编码: 目标编码（如"gbk"、"utf-8"）
//
// 返回值：
//   - string: 转换后的文本
//
// 示例：
//
//	E文本编码转换("测试", "gbk", "utf-8")
func E文本编码转换(str interface{}, 来源编码 string, 目标编码 string) string {
	if 来源编码 == "" {
		来源编码 = E编码_检测(E到文本(str))
		// 如果编码是一致的那么就不需要转换了
		if 来源编码 == 目标编码 {
			return E到文本(str)
		}
		if !(来源编码 == "gbk" || 来源编码 == "utf-8") {
			return E到文本(str)
		}
	}
	srcDecoder := mahonia.NewDecoder(来源编码)
	desDecoder := mahonia.NewDecoder(目标编码)
	resStr := srcDecoder.ConvertString(E到文本(str))
	_, resBytes, _ := desDecoder.Translate(E到字节集(resStr), true)
	return E到文本(resBytes)
}

// E编码_是否为gbk 判断字节数据是否为GBK编码。
//
// 参数：
//   - data: 要判断的字节数据
//
// 返回值：
//   - bool: 是GBK编码返回true，否则返回false
func E编码_是否为gbk(data []byte) bool {
	return E编码_检测(data) == "gbk"
}

// E编码_是否为utf8 判断字节数据是否为UTF-8编码。
//
// 参数：
//   - data: 要判断的字节数据
//
// 返回值：
//   - bool: 是UTF-8编码返回true，否则返回false
func E编码_是否为utf8(data []byte) bool {
	return E编码_检测(data) == "utf-8"
}

// E编码_utf8到gbk 将UTF-8编码的文本转换为GBK编码。
//
// 参数：
//   - str: UTF-8编码的文本
//
// 返回值：
//   - string: GBK编码的文本
func E编码_utf8到gbk(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	return string(gbkData)
}

// E编码_gbk到utf8 将GBK编码的文本转换为UTF-8编码。
//
// 参数：
//   - str: GBK编码的文本
//
// 返回值：
//   - string: UTF-8编码的文本
func E编码_gbk到utf8(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(str))
	return string(gbkData)
}

// E编码_检测 自动检测文本的编码类型。
//
// 参数：
//   - s: 要检测的文本或字节数据
//
// 返回值：
//   - string: 检测到的编码类型（如"utf-8"、"gbk"）
func E编码_检测(s interface{}) string {
	return chardet.Mostlike(E到字节集(s))
}
