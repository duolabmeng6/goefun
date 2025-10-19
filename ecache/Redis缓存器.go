package ecache

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/duolabmeng6/goefun/ecore"
    "github.com/gomodule/redigo/redis"
    "sync"
    "time"
)

// Redis缓存实现
type Redis缓存器 struct {
    互斥锁   sync.Mutex
    redis *redis.Pool
}

// NewRedis缓存器 创建基于 Redis 的缓存器
// link 为 Redis 地址（如 127.0.0.1:6379），password 为密码，db 为数据库编号（0-15）
func NewRedis缓存器(link string, password string, db int) *Redis缓存器 {
    redisPool := &redis.Pool{
        MaxIdle:     100,
        MaxActive:   0,
        IdleTimeout: 240 * time.Second,
        Wait:        true,
        Dial: func() (redis.Conn, error) {
            con, err := redis.Dial("tcp", link,
                redis.DialPassword(password),
                redis.DialDatabase(int(db)),
                redis.DialConnectTimeout(240*time.Second),
                redis.DialReadTimeout(240*time.Second),
                redis.DialWriteTimeout(240*time.Second))
            if err != nil {
                return nil, err
            }
            return con, nil
        },
    }

    return &Redis缓存器{
        互斥锁:   sync.Mutex{},
        redis: redisPool,
    }
}

// Set 写入缓存值到 Redis
// 使用 EX 秒作为过期时间。
func (fc *Redis缓存器) Set(key string, value interface{}, 倒计时秒数 int64) error {
    fc.互斥锁.Lock()
    defer fc.互斥锁.Unlock()
    缓存数据 := &缓存数据{
        DataType: fmt.Sprintf("%T", value),
        Data:     value,
        EndTime:  ecore.E取时间戳() + 倒计时秒数,
    }
    内容, _ := json.Marshal(缓存数据)
    // 保存到数据库中
    conn := fc.redis.Get()
    defer conn.Close()
    _, err := conn.Do("SET", key, string(内容), "EX", 倒计时秒数)
    if err != nil {
        return err
    }

    return nil
}

// Get 读取缓存值
// 若缓存不存在或已过期，返回错误。
func (fc *Redis缓存器) Get(key string) (interface{}, error) {
    fc.互斥锁.Lock()
    defer fc.互斥锁.Unlock()

    // 从数据库中获取
    conn := fc.redis.Get()
    defer conn.Close()
    内容, err := redis.Bytes(conn.Do("GET", key))
    if err != nil {
        return nil, err
    }

    缓存数据 := &缓存数据{}
    err = json.Unmarshal(内容, 缓存数据)
    if err != nil {
        return nil, err
    }
    //检查 缓存数据.EndTime 是否过期
    if 缓存数据.EndTime < ecore.E取时间戳() {
        // 删除数据库中的key
        _, _ = conn.Do("DEL", key)

        return nil, errors.New("缓存已过期")
    }

    return 缓存数据.Data, nil
}

// Del 删除指定 key 的缓存
func (fc *Redis缓存器) Del(key string) error {
    fc.互斥锁.Lock()
    defer fc.互斥锁.Unlock()
    // 删除数据库中的key
    conn := fc.redis.Get()
    defer conn.Close()
    _, err := conn.Do("DEL", key)
    if err != nil {
        return err
    }

    return nil
}
