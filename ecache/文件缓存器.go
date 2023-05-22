package ecache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"sync"
)

// 文件缓存实现
type 文件缓存器 struct {
	互斥锁  sync.Mutex
	存储路径 string
}

func New文件缓存器(存储路径 string) *文件缓存器 {
	return &文件缓存器{
		互斥锁:  sync.Mutex{},
		存储路径: 存储路径,
	}
}

func (fc *文件缓存器) Set(key string, value interface{}, 倒计时秒数 int64) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	文件路径 := fc.存储路径 + key
	缓存数据 := &缓存数据{
		DataType: fmt.Sprintf("%T", value),
		Data:     value,
		EndTime:  ecore.E取时间戳() + 倒计时秒数,
	}
	// 缓存数据 转换为json保存
	内容, err := json.Marshal(缓存数据)
	if err != nil {
		return err
	}
	err = ecore.E写到文件(文件路径, 内容)
	if err != nil {
		return err
	}
	return nil
}

func (fc *文件缓存器) Get(key string) (interface{}, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	文件路径 := fc.存储路径 + key
	内容 := ecore.E读入文件(文件路径)
	if len(内容) == 0 {
		return nil, errors.New("缓存不存在")
	}
	缓存数据 := &缓存数据{}
	err := json.Unmarshal(内容, 缓存数据)
	if err != nil {
		return nil, err
	}
	//检查 缓存数据.EndTime 是否过期
	if 缓存数据.EndTime < ecore.E取时间戳() {
		//log.Fatalln("缓存已过期", 文件路径)
		_ = ecore.E删除文件(文件路径)
		return nil, errors.New(("缓存已过期"))
	}

	return 缓存数据.Data, nil
}

func (fc *文件缓存器) Del(key string) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	文件路径 := fc.存储路径 + key

	return ecore.E删除文件(文件路径)
}
