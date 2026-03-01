package ecore

import (
    "fmt"
    "time"
)

// E时间统计类 用于统计代码执行时间的类
type E时间统计类 struct {
    t time.Time
}

// New时间统计类 创建一个新的时间统计对象
//
// 返回值：
//   *E时间统计类 - 时间统计对象指针
func New时间统计类() *E时间统计类 {
    t := new(E时间统计类)
    t.E开始()
    return t
}

// E开始 开始计时
func (this *E时间统计类) E开始() {
    this.t = time.Now()
}

// E取毫秒 获取从开始到当前的毫秒数
//
// 返回值：
//   string - 毫秒数文本
func (this *E时间统计类) E取毫秒() string {
    return fmt.Sprintf("%d", time.Since(this.t).Milliseconds())
}

// E取秒 获取从开始到当前的秒数
//
// 返回值：
//   string - 秒数文本（保留3位小数）
func (this *E时间统计类) E取秒() string {
    return fmt.Sprintf("%.3f", float64(time.Since(this.t).Milliseconds())/float64(1000))
}
