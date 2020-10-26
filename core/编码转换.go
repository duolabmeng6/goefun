package E

import (
	"github.com/axgle/mahonia"
	"github.com/duolabmeng6/goefun/src/chardet"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gurl"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//调用格式： 〈文本型〉 Base64编码 （字节集 编码数据，［文本型 编码表］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：Base64Encode
//将数据编码到Base64。本命令为初级命令。
//参数<1>的名称为“编码数据”，类型为“字节集（bin）”。要编码的字节集数据。
//参数<2>的名称为“编码表”，类型为“文本型（text）”，可以被省略。除特殊情况下，不建议使用本参数。如果使用本参数，那么编码表长度必须为64位，否则会编码失败。默认编码表：“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”。
//
//操作系统需求： Windows
func Base64编码(data []byte) string {
	return gbase64.EncodeToString(data)
}

//调用格式： 〈字节集〉 Base64解码 （文本型 解码内容，［文本型 编码表］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：Base64Decode
//解码Base64文本到数据。本命令为初级命令。
//参数<1>的名称为“解码内容”，类型为“文本型（text）”。要解码的文本数据。
//参数<2>的名称为“编码表”，类型为“文本型（text）”，可以被省略。除特殊情况下，不建议使用本参数。如果使用本参数，那么编码表长度必须为64位，否则会解码失败。默认编码表：“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”。
//
//操作系统需求： Windows
func Base64解码(data string) string {
	str, _ := gbase64.DecodeToString(data)
	return str
}

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
func URL编码(str string) string {
	return gurl.Encode(str)
}

//调用格式： 〈文本型〉 URL解码 （文本型 解码文本，［文本型 编码格式］） - E2EE互联网服务器套件2.2.3->文本处理
//英文名称：URLDecode
//解码URL内容。本命令为初级命令。
//参数<1>的名称为“解码文本”，类型为“文本型（text）”。要进行URL编码的文本内容。
//参数<2>的名称为“编码格式”，类型为“文本型（text）”，可以被省略。指定编码格式。可使用“#文本编码格式_”开头的常量指定。如果为空则默认为原始的编码。
//
//操作系统需求： Windows
func URL解码(str string) string {
	s, _ := gurl.Decode(str)
	return s
}

//component -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment. See http://php.net/manual/en/function.parse-url.php.
func URL解析(str string,component int) map[string]string {
	s, _ := gurl.ParseURL(str,component)
	return s
}


//E文本编码转换("测试一下","gbk","utf-8")
func E文本编码转换(str interface{}, 来源编码 string, 目标编码 string) string {
	if 来源编码 == "" {
		来源编码 = E编码_检测(E到文本(str))
		//如果编码是一致的那么就不需要转换了
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

func E编码_是否为gbk(data []byte) bool {
	return E编码_检测(data) == "gbk"
}
func E编码_是否为utf8(data []byte) bool {
	return E编码_检测(data) == "utf-8"
}

func E编码_utf8到gbk(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str)) //使用官方库将utf-8转换为gbk
	return string(gbkData)
}

func E编码_gbk到utf8(str string) string {
	gbkData, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(str))
	return string(gbkData)
}

func E编码_检测(s interface{}) string {
	return chardet.Mostlike(E到字节集(s))
}
