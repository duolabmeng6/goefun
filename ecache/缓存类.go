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

// New缓存类 创建一个缓存操作实例，需传入具体的缓存适配器实现。
// 通过该实例可统一操作不同的缓存后端（内存、文件、数据库、Redis 等）。
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

// Set 设置缓存值
// 根据 key 存入任意类型的值，并指定倒计时秒数作为过期时间。
func (c *E缓存类) Set(key string, value interface{}, 倒计时秒数 int64) error {
    return c.缓存.Set(key, value, 倒计时秒数)
}

// Get 获取缓存值
// 返回存储的原始值，若不存在或已过期则返回 nil。
func (c *E缓存类) Get(key string) interface{} {
    value, err := c.缓存.Get(key)
    if err != nil {
        return nil
    }
    return value
}

// GetString 以字符串形式获取缓存值
// 当 key 不存在或已过期时返回空字符串。
func (c *E缓存类) GetString(key string) string {
    value, err := c.缓存.Get(key)
    if err != nil {
        return ""
    }
    return value.(string)
}

// GetInt 以整数形式获取缓存值
// 当 key 不存在或已过期时返回 0。
func (c *E缓存类) GetInt(key string) int64 {
    value, err := c.缓存.Get(key)
    if err != nil {
        return 0
    }
    return ecore.E到整数(value)
}

// GetBool 以布尔值形式获取缓存值
// 当 key 不存在或已过期时返回 false。
func (c *E缓存类) GetBool(key string) bool {
    value, err := c.缓存.Get(key)
    if err != nil {
        return false
    }
    return value.(bool)
}

// GetFloat 以浮点数形式获取缓存值
// 当 key 不存在或已过期时返回 0。
func (c *E缓存类) GetFloat(key string) float64 {
    value, err := c.缓存.Get(key)
    if err != nil {
        return 0
    }
    return value.(float64)
}

// GetBytes 以字节切片形式获取缓存值
// 若缓存中存储的是 Base64 文本，会自动解码为原始字节切片；当 key 不存在或已过期时返回 nil。
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

// Del 删除缓存
// 根据 key 删除对应的缓存项。
func (c *E缓存类) Del(key string) error {
    return c.缓存.Del(key)
}
