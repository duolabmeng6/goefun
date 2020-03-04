package coreUtil

import (
	"github.com/duolabmeng6/goefun/src/sizedwaitgroup"
)

type E线程池 struct {
	swg sizedwaitgroup.SizedWaitGroup
}

func New线程池(最大任务数量 int) *E线程池 {
	this := new(E线程池)
	this.E初始化(最大任务数量)
	return this
}

func (this *E线程池) E初始化(最大任务数量 int) {
	this.swg = sizedwaitgroup.New(最大任务数量)
}

func (this *E线程池) E加入任务() {
	this.swg.Add()
}
func (this *E线程池) E等待任务完成() {
	this.swg.Wait()
}
func (this *E线程池) E完成() {
	this.swg.Done()
}
