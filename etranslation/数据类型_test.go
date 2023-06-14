package etranslation

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"html"
	"testing"
)

func Test内容检测(t *testing.T) {
	内容 := "こんにちは、私はかわいいです"
	println(E内容语言类型检测(内容))
	内容 = "大家好大家好大家好大家好"
	println(E内容语言类型检测(内容))
	内容 = "大家好 我是小可愛 今天我10歲了"
	println(E内容语言类型检测(内容))
}

func Test内容2(t *testing.T) {
	str := "Hello, everyone. I&#39;m a little cute. I&#39;m 10 years old today."
	fmt.Println(html.UnescapeString(str))
}

func TestNew翻译(t *testing.T) {
	ecore.E加载环境变量_从文件("./.env")

	翻译接口 := New翻译()
	翻译接口.E注册服务("百度翻译", New百度翻译(ecore.Env("appid", ""), ecore.Env("secret", "")))
	翻译接口.E注册服务("系统翻译", New系统翻译())
	翻译接口.E注册服务("彩云小译", New彩云小译(ecore.Env("cyxy_token", "")))
	翻译接口.E注册服务("有道翻译", New有道翻译(ecore.Env("yd_app_id", ""), ecore.Env("yd_SecretAccessKey", "")))
	翻译接口.E注册服务("火山翻译", New火山翻译(ecore.Env("hs_AccessKeyID", ""), ecore.Env("hs_SecretAccessKey", "")))
	翻译接口.E注册服务("腾讯翻译", New腾讯翻译(ecore.Env("tx_secretId", ""), ecore.Env("tx_secretKey", "")))
	翻译接口.E注册服务("阿里云翻译", New阿里云翻译(ecore.Env("access_key_id", ""), ecore.Env("access_key_secret", "")))
	翻译接口.E注册服务("必应免费翻译", New必应免费翻译())
	翻译接口.E注册服务("搜狗免费翻译", New搜狗免费翻译())
	翻译接口.E注册服务("爱词霸免费翻译", New爱词霸免费翻译())
	翻译接口.E注册服务("DeepL免费翻译", NewDeepL免费翻译())
	翻译接口.E注册服务("阿里云免费翻译", New阿里云免费翻译())
	//
	//翻译模块列表 := 翻译接口.E列出翻译模块和初始化参数()
	//for _, 模块 := range 翻译模块列表 {
	//	println("模块:", 模块.Name, "模块参数:", 模块.Params)
	//}

	//
	翻译结果, err := 翻译接口.E取翻译模块("阿里云免费翻译").E翻译("大家好 我是小可愛 今天我10歲了", "zh_CN", "en_US")
	fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("阿里云免费翻译").E翻译("大家好 我是小可愛 今天我10歲了", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("DeepL免费翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("DeepL免费翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("爱词霸免费翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("爱词霸免费翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err = 翻译接口.E取翻译模块("搜狗免费翻译").E翻译("苹果", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("搜狗免费翻译").E翻译("苹果", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("必应免费翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("必应免费翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("阿里云翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("阿里云翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("腾讯翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("腾讯翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("火山翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("火山翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("有道翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("有道翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("彩云小译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("彩云小译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//翻译结果, err := 翻译接口.E取翻译模块("系统翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("系统翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//
	//翻译结果, err = 翻译接口.E取翻译模块("百度翻译").E翻译("hello", "zh_CN", "en_US")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)
	//翻译结果, err = 翻译接口.E取翻译模块("百度翻译").E翻译("hello", "en_US", "zh_CN")
	//fmt.Println("翻译结果:", 翻译结果, "翻译出错:", err)

	//语言列表 := New语言列表()

	//ecore.E调试输出(语言列表.E取全部名称())

	//翻译模块列表 := 翻译接口.E列出翻译模块()
	//for _, 模块名称 := range 翻译模块列表 {
	//	翻译结果, err := 翻译接口.E取翻译模块(模块名称).E翻译("hello", "auto", "auto")
	//	if err != nil {
	//		fmt.Printf("模块 %s 翻译出错: %v\n", 模块名称, err)
	//	} else {
	//		fmt.Printf("模块 %s 翻译结果: %s\n", 模块名称, 翻译结果)
	//	}
	//}
}
