// Package useragentparser Description: 用户代理解析类,用于解析用户代理字符串。
package useragentparser

import (
	"github.com/mileusna/useragent"
)

// E用户代理类型 定义了各种操作系统、浏览器和机器人的常量。
type E用户代理类型 struct{}

// 操作系统常量
const (
	EWindows系统 = "Windows系统"
	E安卓系统      = "安卓系统"
	EMacOS系统   = "MacOS系统"
	EIOS系统     = "IOS系统"
	ELinux系统   = "Linux系统"
	E谷歌系统      = "谷歌系统"
)

// 浏览器常量
const (
	EOpera浏览器     = "Opera浏览器"
	EOperaMini浏览器 = "OperaMini浏览器"
	E谷歌浏览器        = "谷歌浏览器"
	E火狐浏览器        = "火狐浏览器"
	EIE浏览器        = "IE浏览器"
	ESafari浏览器    = "Safari浏览器"
	EEdge浏览器      = "Edge浏览器"
)

// 机器人常量
const (
	E谷歌机器人 = "谷歌机器人"
	E推特机器人 = "推特机器人"
	E脸书机器人 = "脸书机器人"
)

// E用户代理版本号类 表示版本号结构。
type E用户代理版本号类 useragent.VersionNo

// E用户代理解析类 是核心解析类,包含了解析后的用户代理信息。
type E用户代理解析类 useragent.UserAgent

// E解析 对用户代理（useragent）字符串进行解析,返回解析结果对象。
func E解析(文本 string) E用户代理解析类 {
	return E用户代理解析类(useragent.Parse(文本))
}

// E为安卓系统 如果系统名称为安卓则返回真。
func (u E用户代理解析类) E为安卓系统() bool {
	return useragent.UserAgent(u).IsAndroid()
}

// E为谷歌系统 如果系统名称为谷歌系统则返回真。
func (u E用户代理解析类) E为谷歌系统() bool {
	return useragent.UserAgent(u).IsChromeOS()
}

// E为IOS系统 如果系统名称为IOS系统则返回真。
func (u E用户代理解析类) E为IOS系统() bool {
	return useragent.UserAgent(u).IsIOS()
}

// E为MacOS系统 如果系统名称为MacOS系统则返回真。
func (u E用户代理解析类) E为MacOS系统() bool {
	return useragent.UserAgent(u).IsMacOS()
}

// E为Linux系统 如果系统名称为Linux系统则返回真。
func (u E用户代理解析类) E为Linux系统() bool {
	return useragent.UserAgent(u).IsLinux()
}

// E为Windows系统 如果系统名称为Windows系统则返回真。
func (u E用户代理解析类) E为Windows系统() bool {
	return useragent.UserAgent(u).IsWindows()
}

// E为谷歌浏览器 如果名称为谷歌浏览器则返回真。
func (u E用户代理解析类) E为谷歌浏览器() bool {
	return useragent.UserAgent(u).IsChrome()
}

// E为火狐浏览器 如果名称为火狐浏览器则返回真。
func (u E用户代理解析类) E为火狐浏览器() bool {
	return useragent.UserAgent(u).IsFirefox()
}

// E为Edge浏览器 如果名称为Edge浏览器则返回真。
func (u E用户代理解析类) E为Edge浏览器() bool {
	return useragent.UserAgent(u).IsEdge()
}

// E为Opera浏览器 如果名称为Opera浏览器则返回真。
func (u E用户代理解析类) E为Opera浏览器() bool {
	return useragent.UserAgent(u).IsOpera()
}

// E为OperaMini浏览器 如果名称为OperaMini浏览器则返回真。
func (u E用户代理解析类) E为OperaMini浏览器() bool {
	return useragent.UserAgent(u).IsOperaMini()
}

// E为Safari浏览器 如果名称为Safari浏览器则返回真。
func (u E用户代理解析类) E为Safari浏览器() bool {
	return useragent.UserAgent(u).IsSafari()
}

// E为IE浏览器 如果名称为IE浏览器则返回真。
func (u E用户代理解析类) E为IE浏览器() bool {
	return useragent.UserAgent(u).IsInternetExplorer()
}

// E为脸书机器人 如果名称为脸书机器人则返回真。
func (u E用户代理解析类) E为脸书机器人() bool {
	return useragent.UserAgent(u).IsFacebookbot()
}

// E为谷歌机器人 如果名称为谷歌机器人则返回真。
func (u E用户代理解析类) E为谷歌机器人() bool {
	return useragent.UserAgent(u).IsGooglebot()
}

// E为推特机器人 如果名称为推特机器人则返回真。
func (u E用户代理解析类) E为推特机器人() bool {
	return useragent.UserAgent(u).IsTwitterbot()
}

// E取系统名称 返回系统名称。
func (u E用户代理解析类) E取系统名称() string {
	switch {
	case u.E为安卓系统():
		return E安卓系统
	case u.E为谷歌系统():
		return E谷歌系统
	case u.E为IOS系统():
		return EIOS系统
	case u.E为MacOS系统():
		return EMacOS系统
	case u.E为Windows系统():
		return EWindows系统
	case u.E为Linux系统():
		return ELinux系统
	default:
		return "未知系统"
	}
}

// E取浏览器名称 返回浏览器名称。
func (u E用户代理解析类) E取浏览器名称() string {
	switch {
	case u.E为谷歌浏览器():
		return E谷歌浏览器
	case u.E为火狐浏览器():
		return E火狐浏览器
	case u.E为Edge浏览器():
		return EEdge浏览器
	case u.E为Opera浏览器():
		return EOpera浏览器
	case u.E为OperaMini浏览器():
		return EOperaMini浏览器
	case u.E为Safari浏览器():
		return ESafari浏览器
	case u.E为IE浏览器():
		return EIE浏览器
	case u.E为脸书机器人():
		return E脸书机器人
	case u.E为谷歌机器人():
		return E谷歌机器人
	case u.E为推特机器人():
		return E推特机器人
	default:
		return "未知名称"
	}
}

// E是否解析成功 如果该包无法可靠地确定用户代理,则本方法返回假。但名称、系统名称等字段可能仍有值。
func (u E用户代理解析类) E是否解析成功() bool {
	return !useragent.UserAgent(u).IsUnknown()
}

// E取系统版本号 返回系统版本号,格式为：主版本号.次版本号.修订版本号。
func (u E用户代理解析类) E取系统版本号() string {
	return useragent.UserAgent(u).OSVersionNoFull()
}

// E取系统版本号2 返回系统版本号,格式为：主版本号.次版本号。
func (u E用户代理解析类) E取系统版本号2() string {
	return useragent.UserAgent(u).OSVersionNoShort()
}

// E取版本号 返回版本号,格式为：主版本号.次版本号.修订版本号。
func (u E用户代理解析类) E取版本号() string {
	return useragent.UserAgent(u).VersionNoFull()
}

// E取版本号2 返回版本号,格式为：主版本号.次版本号。
func (u E用户代理解析类) E取版本号2() string {
	return useragent.UserAgent(u).VersionNoShort()
}
