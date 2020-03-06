//定时任务
//提供周期性任务执行功能。使用方式类似于Linux下的Crontab或者Java里的Quartz。本对象线程安全。
package os

import (
	"time"
)

type E时钟 struct {
}

func New时钟() *E时钟 {
	this := new(E时钟)
	return this
}
func (this *E时钟) E创建(fn func() bool, 时钟周期 int64) *E时钟 {
	SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
	return this
}

func (this *E时钟) E创建执行一次(fn func(), 时钟周期 int64) *E时钟 {
	SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
	return this
}

//返回一个函数，无论调用多少次，它只会在指定的间隔内执行一次
func (this *E时钟) E时钟周期函数(fn func(), 时钟周期 int64) func() {
	return Throttle(fn, time.Duration(时钟周期)*time.Millisecond, ThrottleOptions{})
}
