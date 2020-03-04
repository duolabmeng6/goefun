package coreUtil

import "sync"

type E读写共享锁 struct {
	lock sync.RWMutex
}

func (this *E读写共享锁) E开始读() {
	this.lock.RLock()
}
func (this *E读写共享锁) E结束读() {
	this.lock.RUnlock()
}
func (this *E读写共享锁) E开始写() {
	this.lock.Lock()
}
func (this *E读写共享锁) E结束写() {
	this.lock.Unlock()
}
