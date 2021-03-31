package gfUtil

import (
	"github.com/gogf/gf/util/gvalid"
)

//验证规则这里看
//https://goframe.org/pages/viewpage.action?pageId=3670507

func E数据校验_检查Map(等待验证的数据 interface{}, 验证规则 interface{}, 提示消息 ...gvalid.CustomMsg) *gvalid.Error {
	return gvalid.CheckMap(等待验证的数据, 验证规则, 提示消息...)
}
func E数据校验_检查Struct(等待验证的数据 interface{}, 验证规则 interface{}, 提示消息 ...gvalid.CustomMsg) *gvalid.Error {
	return gvalid.CheckStruct(等待验证的数据, 验证规则, 提示消息...)
}
func E数据校验_检查(等待验证的数据 interface{}, 验证规则 string, 提示消息 interface{}, 参数 ...interface{}) *gvalid.Error {
	return gvalid.Check(等待验证的数据, 验证规则, 提示消息, 参数...)
}
