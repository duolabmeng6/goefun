package gfUtil

import (
	E "github.com/duolabmeng6/goefun/core"
	"testing"
)

func Test校验数据(t *testing.T) {
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := []string{
		"passport@required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
		"password2@required|length:6,16#",
	}
	if e := E数据校验_检查Map(params, rules); e != nil {
		E.E调试输出(e.Map())
		E.E调试输出(e.FirstItem())
		E.E调试输出(e.FirstString())
	}
	// May Output:
	// map[required:账号不能为空 length:账号长度应当在6到16之间]
	// passport map[required:账号不能为空 length:账号长度应当在6到16之间]
	// 账号不能为空
}
func Test校验数据_结构体(t *testing.T) {
	type Params struct {
		Page      int    `v:"required|min:1         # page is required"`
		Size      int    `v:"required|between:1,100 # size is required"`
		ProjectId string `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := E数据校验_检查Struct(obj, nil)
	E.E调试输出(err == nil)
}

func Test校验数据_检查(t *testing.T) {
	if e := E数据校验_检查("123456", "length:6,16", nil);  e != nil {
		E.E调试输出(e.String())
	}
	if e := E数据校验_检查("12345", "length:6,16", nil);  e != nil {
		E.E调试输出(e.String())
	}

	if e := E数据校验_检查(5.66, "integer|between:6,16", "请输入一个整数|参数大小不对啊老铁"); e != nil {
		E.E调试输出(e.Map())
	}

	rule := "url|min-length:11"
	msgs := map[string]string{
		"url"        : "请输入正确的URL地址",
		"min-length" : "地址长度至少为:min位",
	}
	if e := E数据校验_检查("https.goframeorg.ggg", rule, msgs); e != nil {
		E.E调试输出(e.Map())
	}

	//使用正则验证
	rule2 := `regex:\d{6,}|\D{6,}|max-length:16`
	if e := E数据校验_检查("1234562222222222222222222222", rule2, "错误了|超度太长");  e != nil {
		E.E调试输出(e.Map())
	}
	if e := E数据校验_检查("abcde6", rule2, "错误了"); e != nil {
		E.E调试输出(e.Map())
	}


}
