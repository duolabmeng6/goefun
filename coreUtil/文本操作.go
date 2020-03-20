package E

import (
	"fmt"
	. "github.com/duolabmeng6/goefun/core"
	"github.com/gogf/gf/util/grand"
	"strings"
	"unicode"
	"unicode/utf8"
)

func E格式化文本(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

//文本截取函数
func StrCut(内容 string, 表达式 string) string {
	args := strings.Split(表达式, "$")
	if len(args) == 2 {
		return 文本_取出中间文本(内容, args[0], args[1])
	}
	if len(args) == 1 {
		return 文本_取出中间文本(内容, args[0], "")
	}
	return ""
}

func E文本_取左边(被查找的文本 string, 欲寻找的文本 string) string {
	return 文本_取出中间文本(被查找的文本, "", 欲寻找的文本)
}
func E文本_取右边(被查找的文本 string, 欲寻找的文本 string) string {
	return 文本_取出中间文本(被查找的文本, 欲寻找的文本, "")
}

//文本取出中间文本
func 文本_取出中间文本(内容 string, 左边文本 string, 右边文本 string) string {
	左边位置 := strings.Index(内容, 左边文本)
	if 左边位置 == -1 {
		return ""
	}
	左边位置 = 左边位置 + len(左边文本)
	内容 = string([]byte(内容)[左边位置:])

	var 右边位置 int
	if 右边文本 == "" {
		右边位置 = len(内容)
	} else {
		右边位置 = strings.Index(内容, 右边文本)
		if 右边位置 == -1 {
			return ""
		}
	}
	内容 = string([]byte(内容)[:右边位置])
	return 内容
}

//子程序名：文本_取随机字母
//取随机26个字母！
//返回值类型：文本型
//参数<1>的名称为“要取的字符数”，类型为“整数型”。注明：要取字符个数；。
//参数<2>的名称为“字母类型”，类型为“整数型”，允许接收空参数数据。注明：可空：默认为小写字母；0=小写字母；1=大写字母；2=大小写字母混合；。
func E文本_取随机字母(要取的字符数 int, 字母类型 int) string {
	var str string
	if 字母类型 == 0 {
		str = "abcdefghijklmnopqrstuvwxyz"
	}
	if 字母类型 == 1 {
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if 字母类型 == 2 {
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	return grand.Str(str, 要取的字符数)
}

func E文本_取随机字母和数字(要取的字符数 int) string {
	return grand.Str("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 要取的字符数)
}

func E文本_取随机数字(要取的字符数 int) string {
	return grand.Str("0123456789", 要取的字符数)
}

func E文本_删左边(欲处理文本 string, 删除长度 int64) string {
	return E取文本右边(欲处理文本, E取文本长度(欲处理文本)-删除长度)
}
func E文本_删右边(欲处理文本 string, 删除长度 int64) string {
	return E取文本左边(欲处理文本, E取文本长度(欲处理文本)-删除长度)
}

func E文本_删中间(欲处理文本 string, 起始位置 int64, 删除长度 int64) string {
	return E取文本左边(欲处理文本, 起始位置) + E文本_删左边(欲处理文本, 起始位置+删除长度)
}

func E文本_取出文本中汉字(s string) string {
	return E文本区分_只取汉子(s)
}

func E文本_逐字分割(s string) []string {
	r := []rune(s)
	strarr := []string{}
	for _, s := range r {
		strarr = append(strarr, string(s))
	}
	return strarr
}

func E文本_颠倒(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func E文本_取随机姓氏() string {
	百家姓 := "赵钱孙李周吴郑王冯陈褚卫蒋沈韩杨朱秦尤许何吕施张孔曹严华金魏陶姜戚谢邹喻柏水窦章云苏潘葛奚范彭郎鲁韦昌马苗凤花方俞任袁柳酆鲍史唐费廉岑薛雷贺倪汤滕殷罗毕郝邬安常乐于时傅皮卞齐康伍余元卜顾孟平黄和穆萧尹姚邵湛汪祁毛禹狄米贝明臧计伏成戴谈宋茅庞熊纪舒屈项祝董梁杜阮蓝闵席季麻强贾路娄危江童颜郭梅盛林刁钟徐邱骆高夏蔡田樊胡凌霍虞万支柯昝管卢莫柯房裘缪干解应宗丁宣贲邓郁单杭洪包诸左石崔吉钮龚程嵇邢滑裴陆荣翁荀羊于惠甄曲家封芮羿储靳汲邴糜松井段富巫乌焦巴弓牧隗山谷车侯宓蓬全郗班仰秋仲伊宫宁仇栾暴甘钭历戎祖武符刘景詹束龙叶幸司韶郜黎蓟溥印宿白怀蒲邰从鄂索咸籍赖卓蔺屠蒙池乔阳郁胥能苍双闻莘党翟谭贡劳逄姬申扶堵冉宰郦雍却璩桑桂濮牛寿通边扈燕冀浦尚农温别庄晏柴瞿阎充慕连茹习宦艾鱼容向古易慎戈廖庾终暨居衡步都耿满弘匡国文寇广禄阙东欧殳沃利蔚越夔隆师巩厍聂晁勾敖融冷訾辛阚那简饶空曾毋沙乜养鞠须丰巢关蒯相查后荆红游竺权逮盍益桓公"
	return grand.Str(百家姓, 1)
}
func E文本_自动补零(s string, len int) string {
	return E格式化文本("%0*d", len, E到整数(s))
}

//unicode的参数含义
//https://www.cnblogs.com/golove/p/3269099.html
//Golang学习 - unicode 包
//https://www.cnblogs.com/golove/p/3273585.html

func E文本_是否为小写字母(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func E文本_是否为大写字母(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
func E文本_是否为字母(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
func E文本_是否为数字(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

func E文本_是否为汉字(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func E文本区分_只取字母(s string) string {

	str := ""
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			str = str + string(r)
		}
	}
	return str
}

func E文本区分_只取数字(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsNumber(r) {
			str = str + string(r)
		}
	}
	return str
}

func E文本区分_只取汉子(s string) string {
	str := ""
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			str = str + string(r)
		}
	}
	return str
}
func E文本区分_只取符号(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsSymbol(r) {
			str = str + string(r)
		}
	}
	return str
}

func E文本_首字母改大写(s string) string {
	if len(s) < 1 {
		return ""
	}
	strArry := []rune(s)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func E取文本字数(value string) int {
	return utf8.RuneCountInString(value)
}

//子程序名：判断文本
//真 找到 假 找不到
//返回值类型：逻辑型
//参数<1>的名称为“与判断的文本”，类型为“文本型”。
//参数<2>的名称为“关键字”，类型为“文本型”，允许接收空参数数据。
//参数<3>的名称为“更多关键字”，类型为“文本型”，允许接收空参数数据，需要接收数组数据。
func E判断文本(内容 string, 关键字 ...interface{}) bool {
	for _, v := range 关键字 {
		str := E到文本(v)
		if strings.Index(内容, str) != -1 {
			return true
		}
	}
	return false
}

func E判断文本s(内容 string, 关键字 ...interface{}) string {
	for _, v := range 关键字 {
		str := E到文本(v)
		if strings.Index(内容, str) != -1 {
			return str
		}
	}
	return ""
}

func E判断文本前缀(s string, 前缀 string) bool {
	return strings.HasPrefix(s, 前缀)
}
func E判断文本后缀(s string, 后缀 string) bool {
	return strings.HasSuffix(s, 后缀)
}
