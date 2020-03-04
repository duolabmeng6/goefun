package coreUtil

import (
	"fmt"
	. "github.com/duolabmeng6/goefun/core"
	"math/rand"
	"time"
)


func E取随机ip() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254))
	return ip
}
