package ecache

import (
	"errors"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"sync"
)

// 内存缓存实现
type 内存缓存器 struct {
	互斥锁  sync.Mutex
	储存容器 map[string]*缓存数据
}

func New内存缓存器() *内存缓存器 {
	return &内存缓存器{
		互斥锁:  sync.Mutex{},
		储存容器: make(map[string]*缓存数据),
	}
}

func (fc *内存缓存器) Set(key string, value interface{}, 倒计时秒数 int64) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	缓存数据 := &缓存数据{
		DataType: fmt.Sprintf("%T", value),
		Data:     value,
		EndTime:  ecore.E取时间戳() + 倒计时秒数,
	}

	fc.储存容器[key] = 缓存数据

	return nil
}

func (fc *内存缓存器) Get(key string) (interface{}, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	缓存数据 := fc.储存容器[key]
	if 缓存数据 == nil {
		return nil, errors.New("缓存不存在")
	}
	//检查 缓存数据.EndTime 是否过期
	if 缓存数据.EndTime < ecore.E取时间戳() {
		delete(fc.储存容器, key)
		return nil, errors.New("缓存已过期")
	}

	return 缓存数据.Data, nil
}

func (fc *内存缓存器) Del(key string) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	delete(fc.储存容器, key)
	return nil
}
