package ecache

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/duolabmeng6/goefun/ecore"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "sync"
)

// Sqlite缓存实现
type Sqlite缓存器 struct {
    互斥锁 sync.Mutex
    db  *gorm.DB
}

type cacheTable struct {
    // 设置 Key 为主键 唯一索引 字段为 `k`
    Key string `gorm:"primaryKey;uniqueIndex;column:k"`
    // 设置 Value 字段为 `v`
    Value   string `gorm:"column:v"`
    EndTime int64
}

// 设置表名
func (cacheTable) TableName() string {
    return "ecache"
}

// NewSqlite缓存器 创建基于 Sqlite 的缓存器
// 缓存数据库路径示例：./cache.db 或 :memory:（内存模式）
func NewSqlite缓存器(缓存数据库路径 string) *Sqlite缓存器 {
    db, err := gorm.Open(sqlite.Open(缓存数据库路径), &gorm.Config{})
    //gorm 关闭日志
    db.Logger = db.Logger.LogMode(0)
    //内存模式
    //db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }
    // 自动迁移 schema
    db.AutoMigrate(&cacheTable{})
    return &Sqlite缓存器{
        互斥锁: sync.Mutex{},
        db:  db,
    }
}

// Set 写入缓存值到 Sqlite 数据库
func (fc *Sqlite缓存器) Set(key string, value interface{}, 倒计时秒数 int64) error {
    fc.互斥锁.Lock()
    defer fc.互斥锁.Unlock()
    缓存数据 := &缓存数据{
        DataType: fmt.Sprintf("%T", value),
        Data:     value,
        EndTime:  ecore.E取时间戳() + 倒计时秒数,
    }
    内容, _ := json.Marshal(缓存数据)
    // 如果key存在则更新
    var cache cacheTable
    fc.db.Where("k = ?", key).First(&cache)
    if cache.Key != "" {
        fc.db.Model(&cache).Update("value", string(内容))
        fc.db.Model(&cache).Update("endTime", 缓存数据.EndTime)
        return nil
    } else {
        fc.db.Create(&cacheTable{Key: key,
            Value:   string(内容),
            EndTime: 缓存数据.EndTime})
    }

    return nil
}

// Get 读取缓存值
// 从数据库查询对应 key 的数据；当缓存不存在或已过期时返回错误。
func (fc *Sqlite缓存器) Get(key string) (interface{}, error) {
    fc.互斥锁.Lock()
    defer fc.互斥锁.Unlock()
    // 查询数据库中的key
    var cache cacheTable
    fc.db.Where("k = ?", key).First(&cache)

    if cache.Key == "" {
        return nil, errors.New("缓存不存在")
    }
    内容 := []byte(cache.Value)
    缓存数据 := &缓存数据{}
    err := json.Unmarshal(内容, 缓存数据)
    if err != nil {
        return nil, err
    }
    //检查 缓存数据.EndTime 是否过期
    if 缓存数据.EndTime < ecore.E取时间戳() {
        fc.db.Delete(&cacheTable{}, "key = ?", key)

        return nil, errors.New("缓存已过期")
    }

    return 缓存数据.Data, nil
}

// Del 删除指定 key 的缓存记录
func (fc *Sqlite缓存器) Del(key string) error {
    fc.互斥锁.Lock()
    defer fc.互斥锁.Unlock()
    // 删除数据库中的key
    fc.db.Delete(&cacheTable{}, "k = ?", key)

    return nil
}
