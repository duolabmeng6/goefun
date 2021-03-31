package 数据校验

import "github.com/gogf/gf/util/gvalid"

//验证规则这里看
//https://goframe.org/pages/viewpage.action?pageId=3670507

func E数据校验_检查Map(params interface{}, rules interface{}, messages ...gvalid.CustomMsg) *gvalid.Error {
	return gvalid.CheckMap(params, rules, messages...)
}
func E数据校验_检查Struct(params interface{}, rules interface{}, messages ...gvalid.CustomMsg) *gvalid.Error {
	return gvalid.CheckStruct(params, rules, messages...)
}
func E数据校验_检查(value interface{}, rules string, messages interface{}, params ...interface{}) *gvalid.Error {
	return gvalid.Check(value, rules, messages, params...)
}
