package efile

import (
	"github.com/duolabmeng6/goefun/ecore"
	"sync"
)

// 文件缓存实现
type 本地文件储存器 struct {
	互斥锁  sync.Mutex
	存储路径 string
}

func New本地文件储存器(存储路径 string) *本地文件储存器 {
	//存储路径 是否为空 为空则使用默认路径 ./storage/
	if 存储路径 == "" {
		存储路径 = "./storage/"
	}
	//检查路径后缀是否有 / 如果没有则加上
	if ecore.E取文本右边(存储路径, 1) != "/" {
		存储路径 = 存储路径 + "/"
	}

	return &本地文件储存器{
		互斥锁:  sync.Mutex{},
		存储路径: 存储路径,
	}
}

// Put
func (fc *本地文件储存器) Put(key string, data interface{}) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	return ecore.E写到文件(fc.存储路径+key, ecore.E到字节集(data))
}

// Get
func (fc *本地文件储存器) Get(key string) ([]byte, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	数据 := ecore.E读入文件(fc.存储路径 + key)
	return 数据, nil
}

// Delete
func (fc *本地文件储存器) Delete(key string) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	return ecore.E删除文件(fc.存储路径 + key)
}

// Move
func (fc *本地文件储存器) Move(key string, tokey string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	err := ecore.E移动文件(fc.存储路径+key, fc.存储路径+tokey)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Copy
func (fc *本地文件储存器) Copy(key string, tokey string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	err := ecore.E复制文件(fc.存储路径+key, fc.存储路径+tokey)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Exists
func (fc *本地文件储存器) Exists(key string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	return ecore.E文件是否存在(fc.存储路径 + key), nil
}

// Size
func (fc *本地文件储存器) Size(key string) (int64, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	return ecore.E取文件尺寸(fc.存储路径 + key), nil
}

// List
func (fc *本地文件储存器) List(key string) ([]string, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 枚举文件
	文件列表 := []string{}
	err := ecore.E文件枚举(fc.存储路径+key, "", &文件列表, false, false)
	if err != nil {
		return nil, err
	}

	return 文件列表, nil
}

// MimeType
func (fc *本地文件储存器) MimeType(key string) (string, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	return ecore.E取文件Mime(fc.存储路径 + key), nil
}
