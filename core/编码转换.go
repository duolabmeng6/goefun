package E

import (
	"fmt"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//调用格式： 〈文本型〉 Base64编码 （字节集 编码数据，［文本型 编码表］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：Base64Encode
//将数据编码到Base64。本命令为初级命令。
//参数<1>的名称为“编码数据”，类型为“字节集（bin）”。要编码的字节集数据。
//参数<2>的名称为“编码表”，类型为“文本型（text）”，可以被省略。除特殊情况下，不建议使用本参数。如果使用本参数，那么编码表长度必须为64位，否则会编码失败。默认编码表：“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”。
//
//操作系统需求： Windows

//调用格式： 〈字节集〉 Base64解码 （文本型 解码内容，［文本型 编码表］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：Base64Decode
//解码Base64文本到数据。本命令为初级命令。
//参数<1>的名称为“解码内容”，类型为“文本型（text）”。要解码的文本数据。
//参数<2>的名称为“编码表”，类型为“文本型（text）”，可以被省略。除特殊情况下，不建议使用本参数。如果使用本参数，那么编码表长度必须为64位，否则会解码失败。默认编码表：“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”。
//
//操作系统需求： Windows

//调用格式： 〈文本型〉 创建文本 （文本型 待格式化文本，［通用型 待格式化参数］，... ） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：CreateString
//通过文本模板建立文本。例如：“我叫{1}”，{1}表示第一个参数。支持转义符 \r\n(换行)、\t(制表)、\'(单引号)、\{、\}(左右大括号)、\\(\本身)。其它时间\将被保留。本命令为中级命令。命令参数表中最后一个参数可以被重复添加。
//参数<1>的名称为“待格式化文本”，类型为“文本型（text）”。格式为大括号括起来的参数位置，例如：“我叫{1}”，{1} 表示第一个参数，以此类推。两个单引号“``”表示双引号，对应位置参数为空或者不支持时则填充为空文本。支持转义符 \r\n(换行)、\t(制表)、\'(单引号)、\{、\}(左右大括号)。其它时间\将被保留。
//参数<2>的名称为“待格式化参数”，类型为“通用型（all）”，可以被省略。存取键值表等通用对象支持的格式的参数，不支持则会填充为空文本。本参数可以为空。
//
//操作系统需求： Windows

//调用格式： 〈文本型〉 HTML关键字编码 （文本型 HTML文本） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：HtmlspecialcharsEncode
//将HTML中的关键字 &、"、'、<、> 转换成转义符 &amp;、&quot;、&apos;、&lt;、&gt;。本命令为初级命令。
//参数<1>的名称为“HTML文本”，类型为“文本型（text）”。待编码的HTML文本。
//
//操作系统需求： Windows

//调用格式： 〈文本型〉 HTML关键字解码 （文本型 HTML文本） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：HtmlspecialcharsDecode
//将转义的HTML中的转义符 &amp;、&quot;、&apos;、&lt;、&gt; 恢复成关键字 &、"、'、<、>。本命令为初级命令。
//参数<1>的名称为“HTML文本”，类型为“文本型（text）”。待解码的HTML文本。
//
//操作系统需求： Windows

//调用格式： 〈文本型〉 URL编码 （文本型 编码文本，［文本型 编码格式］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：URLEncode
//编码URL内容。本命令为初级命令。
//参数<1>的名称为“编码文本”，类型为“文本型（text）”。要进行URL编码的文本内容。
//参数<2>的名称为“编码格式”，类型为“文本型（text）”，可以被省略。指定编码格式。可使用“#文本编码格式_”开头的常量指定。如果为空则默认原始编码。
//
//操作系统需求： Windows

//调用格式： 〈文本型〉 URL解码 （文本型 解码文本，［文本型 编码格式］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：URLDecode
//解码URL内容。本命令为初级命令。
//参数<1>的名称为“解码文本”，类型为“文本型（text）”。要进行URL编码的文本内容。
//参数<2>的名称为“编码格式”，类型为“文本型（text）”，可以被省略。指定编码格式。可使用“#文本编码格式_”开头的常量指定。如果为空则默认为原始的编码。
//
//操作系统需求： Windows

//E文本编码转换("测试一下","gbk","utf-8")
func E文本编码转换(str interface{}, 来源编码 string, 目标编码 string) string {
	if 来源编码 == "" {
		if E编码_是否为gbk(E到字节集(str)) {
			来源编码 = "gbk"
		}
		if E编码_是否为utf8(E到字节集(str)) {
			来源编码 = "utf-8"
		}
		//如果编码是一致的那么就不需要转换了
		if 来源编码 == 目标编码 {
			return E到文本(str)
		}
	}
	srcDecoder := mahonia.NewDecoder(来源编码)
	desDecoder := mahonia.NewDecoder(目标编码)
	resStr := srcDecoder.ConvertString(E到文本(str))
	_, resBytes, _ := desDecoder.Translate(E到字节集(resStr), true)
	return E到文本(resBytes)
}

func E编码_是否为gbk(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
func preNUm(data byte) int {
	str := fmt.Sprintf("%b", data)
	var i int = 0
	for i < len(str) {
		if str[i] != '1' {
			break
		}
		i++
	}
	return i
}
func E编码_是否为utf8(data []byte) bool {
	for i := 0; i < len(data); {
		if data[i]&0x80 == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if data[i]&0xc0 != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func E编码_utf8到gbk(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str)) //使用官方库将utf-8转换为gbk
	return string(gbkData)
}

func E编码_gbk到utf8(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(str))
	return string(gbkData)
}
