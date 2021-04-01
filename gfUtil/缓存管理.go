package gfUtil

import (
	"time"

	"github.com/gogf/gf/os/gcache"
)

//gcache模块默认提供的是一个高速的内存缓存，操作效率非常高效，CPU性能损耗在ns纳秒级别。
//https://goframe.org/pages/viewpage.action?pageId=1114311
type E缓存类 struct {
	c *gcache.Cache
}

func New缓存类() *E缓存类 {
	x := new(E缓存类)
	x.c = gcache.New()
	return x
}

func (this *E缓存类) E设置(key interface{}, value interface{}, duration time.Duration) error {
	return this.c.Set(key, value, duration)
}
func (this *E缓存类) E获取(key interface{}) (interface{}, error) {
	return this.c.Get(key)
}
func (this *E缓存类) E大小() (size int, err error) {
	return this.c.Size()
}

func (this *E缓存类) E是否存在(key interface{}) (bool, error) {
	return this.c.Contains(key)
}

func (this *E缓存类) E删除(keys ...interface{}) (value interface{}, err error) {
	return this.c.Remove(keys...)
}
func (this *E缓存类) E关闭() error {
	return this.c.Close()
}

func (this *E缓存类) E清除() error {
	return this.c.Clear()
}

// Values以slice的形式返回缓存中的所有值。
func (this *E缓存类) E获取所有值() ([]interface{}, error) {
	return this.c.Values()
}

// Keys以slice的形式返回缓存中的所有键。
func (this *E缓存类) E获取所有键() ([]interface{}, error) {
	return this.c.Keys()
}
