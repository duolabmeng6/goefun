package chardet

import (
	"bytes"
	"errors"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
)

// 提供的编码格式字符串未知时，会返回本错误
var ErrUnknown = errors.New("unknown codec")

// Codec映射编码格式字符串到对应的编解码器，因此不需在导入encoding等包
var Codec = map[string]encoding.Encoding{
	"utf-8":    encoding.Nop,
	"utf-16be": unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
	"utf-16le": unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
	//"utf-32be":    UTF32BE,
	//"utf-32le":    UTF32LE,
	"hz-gb2312":   simplifiedchinese.HZGB2312,
	"gbk":         simplifiedchinese.GBK,
	"big5":        traditionalchinese.Big5,
	"gb18030":     simplifiedchinese.GB18030,
	"euc-kr":      korean.EUCKR,
	"euc-jp":      japanese.EUCJP,
	"iso-2022-jp": japanese.ISO2022JP,
	"shift-jis":   japanese.ShiftJIS,
}

// 函数返回一个Reader接口，该接口将从r读取的数据解码后返回；
// codec参数指定编码格式，data为从r读取以检测编码格式的数据；
// 函数会首先解码data，以便返回完整的解码后文本，会自动处理BOM。
func NewReader(r io.Reader, codec string, data []byte) (io.Reader, error) {
	switch codec {
	case "utf-16be":
		if len(data) >= 2 && string(data[:2]) == "\xFE\xFF" {
			data = data[2:]
		}
	case "utf-16le":
		if len(data) >= 2 && string(data[:2]) == "\xFF\xFE" {
			data = data[2:]
		}
	case "utf-8":
		if len(data) >= 3 && string(data[:3]) == "\xEF\xBB\xBF" {
			data = data[3:]
		}
	case "utf-32be":
		if len(data) >= 4 && string(data[:4]) == "\x00\x00\xFE\xFF" {
			data = data[4:]
		}
	case "utf-32le":
		if len(data) >= 4 && string(data[:4]) == "\xFF\xFE\x00\x00" {
			data = data[4:]
		}
	case "gb18030":
		if len(data) >= 4 && string(data[:4]) == "\x84\x31\x95\x33" {
			data = data[4:]
		}
	}
	if c, ok := Codec[codec]; ok {
		return transform.NewReader(io.MultiReader(bytes.NewReader(data), r), c.NewDecoder()), nil
	} else {
		return nil, ErrUnknown
	}
}

// 函数返回一个Writer接口，该接口将提供的数据编码后写入w；
// codec参数指定编码格式，如果bom为真，会在w开始处写入BOM标识。
func NeWriter(w io.Writer, codec string, bom bool) (io.Writer, error) {
	if bom {
		switch codec {
		case "utf-16be":
			w.Write([]byte("\xFE\xFF"))
		case "utf-16le":
			w.Write([]byte("\xFF\xFE"))
		case "utf-8":
			w.Write([]byte("\xEF\xBB\xBF"))
		case "utf-32be":
			w.Write([]byte("\x00\x00\xFE\xFF"))
		case "utf-32le":
			w.Write([]byte("\xFF\xFE\x00\x00"))
		case "gb18030":
			w.Write([]byte("\x84\x31\x95\x33"))
		}
	}
	if c, ok := Codec[codec]; ok {
		return transform.NewWriter(w, c.NewEncoder()), nil
	} else {
		return nil, ErrUnknown
	}
}
