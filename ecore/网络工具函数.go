package ecore

import (
    "fmt"
    "math/rand"
    "time"
)

// E取随机ip 生成一个随机的IPv4地址
//
// 返回值：
//   string - 随机IP地址
func E取随机ip() string {
    rand.Seed(time.Now().Unix())
    ip := fmt.Sprintf("%d.%d.%d.%d", E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254))
    return ip
}
