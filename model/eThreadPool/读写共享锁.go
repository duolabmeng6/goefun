package eThreadPool

import "sync"

type E读写共享锁 struct {
	lock sync.RWMutex
}

// New读写共享锁 创建一个读写共享锁
func New读写共享锁() *E读写共享锁 {
	var lock sync.RWMutex
	return &E读写共享锁{lock: lock}
}

// E开始读
func (this *E读写共享锁) E开始读() {
	this.lock.RLock()
}

// E结束读
func (this *E读写共享锁) E结束读() {
	this.lock.RUnlock()
}

// E开始写
func (this *E读写共享锁) E开始写() {
	this.lock.Lock()
}

// E结束写
func (this *E读写共享锁) E结束写() {
	this.lock.Unlock()
}
