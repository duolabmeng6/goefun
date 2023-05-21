package ecore

import (
	"fmt"
	"time"
)

type E时间统计类 struct {
	t time.Time
}

func New时间统计类() *E时间统计类 {
	t := new(E时间统计类)
	t.E开始()
	return t
}
func (this *E时间统计类) E开始() {
	this.t = time.Now()
}

func (this *E时间统计类) E取毫秒() string {
	return fmt.Sprintf("%d", time.Since(this.t).Milliseconds())
}
func (this *E时间统计类) E取秒() string {
	return fmt.Sprintf("%.3f", float64(time.Since(this.t).Milliseconds())/float64(1000))
}
