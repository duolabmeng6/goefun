// Package 提供统一的缓存接口
package ecache

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/gogf/gf/v2/encoding/gbase64"
)

// 缓存适配器
type E缓存类 struct {
	缓存 缓存接口
}

func New缓存类(缓存 缓存接口) *E缓存类 {
	return &E缓存类{
		缓存: 缓存,
	}
}

// E设置值
func (c *E缓存类) E设置值(key string, value interface{}, 倒计时秒数 int64) error {
	return c.缓存.Set(key, value, 倒计时秒数)
}

// E取值
func (c *E缓存类) E取值(key string) interface{} {
	value, _ := c.缓存.Get(key)
	return value
}

// E删除
func (c *E缓存类) E删除(key string) error {
	return c.缓存.Del(key)
}

// Set
func (c *E缓存类) Set(key string, value interface{}, 倒计时秒数 int64) error {
	return c.缓存.Set(key, value, 倒计时秒数)
}

// Get
func (c *E缓存类) Get(key string) interface{} {
	value, err := c.缓存.Get(key)
	if err != nil {
		return nil
	}
	return value
}

// GetString
func (c *E缓存类) GetString(key string) string {
	value, err := c.缓存.Get(key)
	if err != nil {
		return ""
	}
	return value.(string)
}

// GetInt
func (c *E缓存类) GetInt(key string) int64 {
	value, err := c.缓存.Get(key)
	if err != nil {
		return 0
	}
	return ecore.E到整数(value)
}

// GetBool
func (c *E缓存类) GetBool(key string) bool {
	value, err := c.缓存.Get(key)
	if err != nil {
		return false
	}
	return value.(bool)
}

// GetFloat
func (c *E缓存类) GetFloat(key string) float64 {
	value, err := c.缓存.Get(key)
	if err != nil {
		return 0
	}
	return value.(float64)
}

// GetBytes
func (c *E缓存类) GetBytes(key string) []byte {
	value, err := c.缓存.Get(key)
	if err != nil {
		return nil
	}
	v, err := gbase64.DecodeToString(value.(string))
	if err != nil {
		return nil
	}
	return []byte(v)
}

// Del
func (c *E缓存类) Del(key string) error {
	return c.缓存.Del(key)
}
