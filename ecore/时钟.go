package ecore

//定时任务
//提供周期性任务执行功能。使用方式类似于Linux下的Crontab或者Java里的Quartz。本对象线程安全。
import (
    "github.com/duolabmeng6/goefun/src/cable"
    "time"
)

// E时钟 时钟类，用于创建定时任务和延时执行
type E时钟 struct {
}

// New时钟 创建一个新的时钟对象
//
// 返回值：
//   *E时钟 - 时钟对象指针
func New时钟() *E时钟 {
    this := new(E时钟)
    return this
}

// E创建 创建一个周期性执行的定时器
//
// 参数：
//   fn func() bool - 要执行的函数，返回false停止执行
//   时钟周期 int64 - 执行周期（毫秒）
func (this *E时钟) E创建(fn func() bool, 时钟周期 int64) {
    cable.SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E创建执行一次 创建一个只执行一次的延时定时器
//
// 参数：
//   fn func() - 要执行的函数
//   时钟周期 int64 - 延时时间（毫秒）
func (this *E时钟) E创建执行一次(fn func(), 时钟周期 int64) {
    cable.SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E时钟周期函数 返回一个节流函数，在指定周期内只执行一次
//
// 参数：
//   fn func() - 要执行的函数
//   时钟周期 int64 - 节流周期（毫秒）
//
// 返回值：
//   func() - 节流后的函数
func (this *E时钟) E时钟周期函数(fn func(), 时钟周期 int64) func() {
    return cable.Throttle(fn, time.Duration(时钟周期)*time.Millisecond, cable.ThrottleOptions{})
}

// E时钟_创建 创建一个周期性执行的定时器（全局函数）
//
// 参数：
//   fn func() bool - 要执行的函数，返回false停止执行
//   时钟周期 int64 - 执行周期（毫秒）
func E时钟_创建(fn func() bool, 时钟周期 int64) {
    cable.SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E时钟_创建执行一次 创建一个只执行一次的延时定时器（全局函数）
//
// 参数：
//   fn func() - 要执行的函数
//   时钟周期 int64 - 延时时间（毫秒）
func E时钟_创建执行一次(fn func(), 时钟周期 int64) {
    cable.SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E时钟_创建周期函数 返回一个节流函数，在指定周期内只执行一次（全局函数）
//
// 参数：
//   fn func() - 要执行的函数
//   时钟周期 int64 - 节流周期（毫秒）
//
// 返回值：
//   func() - 节流后的函数
func E时钟_创建周期函数(fn func(), 时钟周期 int64) func() {
    return cable.Throttle(fn, time.Duration(时钟周期)*time.Millisecond, cable.ThrottleOptions{})
}
