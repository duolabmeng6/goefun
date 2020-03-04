package core

import (
	"github.com/gogf/gf/os/gtime"
)

//A simple extension for Time based on PHP's Carbon library. https://github.com/uniplaces/carbon

//到时间
//增减时间
//取时间间隔
//取某月天数
//时间到文本
//取时间部分
//取年份-
//取月份-
//取日-
//取星期几
//取小时
//取分钟
//取秒
//指定时间
//取现行时间
//置现行时间
//取日期
//取时间

type E日期时间型 struct {
	Time *gtime.Time
}

func (this *E日期时间型) E取星期几() int64 {
	return E到整数(this.Time.Format("N"))
}
func (this *E日期时间型) E取月天数() int64 {
	return E到整数(this.Time.Format("t"))
}

func (this *E日期时间型) E取年份() int64 {
	return E到整数(this.Time.Format("Y"))
}
func (this *E日期时间型) E取月份() int64 {
	return E到整数(this.Time.Format("m"))
}
func (this *E日期时间型) E取日() int64 {
	return E到整数(this.Time.Format("d"))
}
func (this *E日期时间型) E取小时() int64 {
	return E到整数(this.Time.Format("H"))
}
func (this *E日期时间型) E取分钟() int64 {
	return E到整数(this.Time.Format("i"))
}
func (this *E日期时间型) E取秒() int64 {
	return E到整数(this.Time.Format("s"))
}
func (this *E日期时间型) E取毫秒() int64 {
	return E到整数(this.Time.Millisecond())
}
func (this *E日期时间型) E取微秒() int64 {
	return E到整数(this.Time.Microsecond())
}
func (this *E日期时间型) E取纳秒() int64 {
	return E到整数(this.Time.Nanosecond())
}
func (this *E日期时间型) E取时间戳() int64 {
	return this.Time.Timestamp()
}
func (this *E日期时间型) E取时间戳毫秒() int64 {
	return this.Time.TimestampMilli()
}
func (this *E日期时间型) E取时间戳微秒() int64 {
	return this.Time.TimestampMicro()
}
func (this *E日期时间型) E取时间戳纳秒() int64 {
	return this.Time.TimestampNano()
}

func (this *E日期时间型) E时间到文本(format string) string {
	if format == "" {
		format = "Y-m-d H:i:s"
	}
	return this.Time.Format(format)
}

func (this *E日期时间型) E增减日期(年 int, 月 int, 日 int) *E日期时间型 {
	this.Time = this.Time.AddDate(年, 月, 日)
	return this
}
func (this *E日期时间型) E增减时间(时 int, 分 int, 秒 int) *E日期时间型 {
	if 时 != 0 {
		this.Time.AddStr(E到文本(时) + "h")
	}
	if 分 != 0 {
		this.Time.AddStr(E到文本(分) + "m")
	}
	if 秒 != 0 {
		this.Time.AddStr(E到文本(秒) + "s")
	}
	return this
}

func (this *E日期时间型) E大于(time *E日期时间型) bool {
	return this.Time.After(time.Time)
}
func (this *E日期时间型) E小于(time *E日期时间型) bool {
	return this.Time.Before(time.Time)
}
func (this *E日期时间型) E等于(time *E日期时间型) bool {
	return this.Time.Equal(time.Time)
}

//到时间
//增减时间
//取时间间隔
//取某月天数
//时间到文本
//取时间部分

func E取现行时间() *E日期时间型 {
	this := new(E日期时间型)
	this.Time = gtime.Now()
	return this
}
func E到时间(s string) *E日期时间型 {
	this := new(E日期时间型)
	if t, err := gtime.StrToTime(s); err == nil {
		this.Time = t
	}
	return this
}
