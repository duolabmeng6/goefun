package ecore

import (
	"testing"
)

func TestE文本编码转换(t *testing.T) {
	type args struct {
		src        string
		oldEncoder string
		newEncoder string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E文本编码转换(tt.args.src, tt.args.oldEncoder, tt.args.newEncoder); got != tt.want {
				t.Errorf("E文本编码转换() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestE文本编码转换utf8(t *testing.T) {

	str := "测试一下，编码问题，6666666666，!@#$%^&*()_+{}|:>?<" //go字符串编码为utf-8
	E调试输出("before convert:", str)                     //打印转换前的字符串
	E调试输出("isUtf8:", E编码_是否为utf8([]byte(str)))        //判断是否是utf-8
	gbkData := E编码_utf8到gbk(str)                      //使用官方库将utf-8转换为gbk
	E调试输出("gbk直接打印会出现乱码:", gbkData)                   //乱码字符串
	E调试输出("isGBK:", E编码_是否为gbk([]byte(gbkData)))      //判断是否是gbk
	utf8Data := E编码_gbk到utf8(gbkData)                 //将gbk再转换为utf-8
	E调试输出("isUtf8:", E编码_是否为utf8([]byte(utf8Data)))   //判断是否是utf-8
	E调试输出("after convert:", utf8Data)                 //打印转换前的字符串

	utf8togbk := E文本编码转换(utf8Data, "utf-8", "gbk")
	E调试输出("utf8 to gbk", utf8togbk)

	gbktoutf8 := E文本编码转换(gbkData, "gbk", "utf-8")
	E调试输出("gbk to utf8", gbktoutf8)

	test1 := E文本编码转换(gbkData, "", "utf-8")
	E调试输出("输入utf8", test1)
	test2 := E文本编码转换(gbktoutf8, "", "utf-8")
	E调试输出("输入gbk", test2)

}

func TestE文本编码转换utf8文件(t *testing.T) {
	str := E到文本(E读入文件("./test.html"))
	E调试输出(E编码_检测(str))

	strgbk := E到文本(E读入文件("./gbk.html"))
	E调试输出(E编码_检测(strgbk))

	str2 := E文本编码转换(strgbk, "", "utf-8")
	//E调试输出(str2)
	//E调试输出(E编码_检测("str1"))

	str3 := E文本编码转换(str2, "", "utf-8")
	//E调试输出(str3)
	E调试输出(E编码_检测(str3))

	str4 := E文本编码转换(str, "", "utf-8")
	//E调试输出(str3)
	E调试输出(E编码_检测(str4))
	E调试输出(str4)

}

func TestBase64解码(t *testing.T) {
	base64string := Base64编码(E到字节集("abcdefg"))
	E调试输出(base64string)
	E调试输出(Base64解码(base64string))
	E调试输出(URL编码("<>?"))
	E调试输出(URL解码("%3C%3E%3F"))
	E调试输出(URL解析("http://user:pass@example.com:8080/path/to/index.html", -1))

}
