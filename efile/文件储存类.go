// Package 提供通用的文件储存类 支持 本地文件储存、阿里云OSS储存、七牛云储存
package efile

// 文件储存类
type 文件储存类 struct {
	文件储存类 map[string]文件储存接口
	当前储存器 string
}

// New文件储存类
// 默认储存器 为 例如 local 为本地文件储存器
func New文件储存类(默认储存器 string) *文件储存类 {
	return &文件储存类{
		文件储存类: make(map[string]文件储存接口),
		当前储存器: 默认储存器,
	}
}

func (c *文件储存类) 设置储存器(储存器名称 string, disk 文件储存接口) *文件储存类 {
	c.文件储存类[储存器名称] = disk
	return c
}

// Disk 获取文件储存类
// 储存器名称 例如 local 为本地文件储存器
func (c *文件储存类) Disk(储存器名称 string) *文件储存类 {
	return &文件储存类{
		文件储存类: c.文件储存类,
		当前储存器: 储存器名称,
	}
}

// getDisk 获取储存器
func (c *文件储存类) getDisk() 文件储存接口 {
	if _, ok := c.文件储存类[c.当前储存器]; !ok {
		panic("没有设置储存器")
	}
	return c.文件储存类[c.当前储存器]
}

// Put 保存文件
func (c *文件储存类) Put(key string, value interface{}) error {
	return c.getDisk().Put(key, value)
}

// Get 获取文件
func (c *文件储存类) Get(key string) ([]byte, error) {
	return c.getDisk().Get(key)
}

// Delete 删除文件
func (c *文件储存类) Delete(key string) error {
	return c.getDisk().Delete(key)
}

// Move 移动文件
func (c *文件储存类) Move(key string, tokey string) (bool, error) {
	return c.getDisk().Move(key, tokey)
}

// Copy 复制文件
func (c *文件储存类) Copy(key string, tokey string) (bool, error) {
	return c.getDisk().Copy(key, tokey)
}

// Exists 文件是否存在
func (c *文件储存类) Exists(key string) (bool, error) {
	return c.getDisk().Exists(key)
}

// Size 文件大小
func (c *文件储存类) Size(key string) (int64, error) {
	return c.getDisk().Size(key)
}

// List 获取文件列表
func (c *文件储存类) List(key string) ([]string, error) {
	return c.getDisk().List(key)
}

// MimeType 获取文件类型
func (c *文件储存类) MimeType(key string) (string, error) {
	return c.getDisk().MimeType(key)
}

// E保存文件
func (c *文件储存类) E保存文件(key string, value interface{}) error {
	return c.getDisk().Put(key, value)
}

// E获取文件
func (c *文件储存类) E获取文件(key string) ([]byte, error) {
	return c.getDisk().Get(key)
}

// E删除文件
func (c *文件储存类) E删除文件(key string) error {
	return c.getDisk().Delete(key)
}

// E移动文件
func (c *文件储存类) E移动文件(key string, tokey string) (bool, error) {
	return c.getDisk().Move(key, tokey)
}

// E复制文件
func (c *文件储存类) E复制文件(key string, tokey string) (bool, error) {
	return c.getDisk().Copy(key, tokey)
}

// E文件是否存在
func (c *文件储存类) E文件是否存在(key string) (bool, error) {
	return c.getDisk().Exists(key)
}

// E文件大小
func (c *文件储存类) E获取文件大小(key string) (int64, error) {
	return c.getDisk().Size(key)
}

// E获取文件列表
func (c *文件储存类) E获取文件列表(key string) ([]string, error) {
	return c.getDisk().List(key)
}

// E获取文件类型
func (c *文件储存类) E获取文件类型(key string) (string, error) {
	return c.getDisk().MimeType(key)
}
