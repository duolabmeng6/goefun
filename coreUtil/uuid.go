package coreUtil

import (
	_ "crypto/md5"
	"github.com/gogf/gf/util/guuid"
)

func E取uuid() string {
	return guuid.New().String()
}
