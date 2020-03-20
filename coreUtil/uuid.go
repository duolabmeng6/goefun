package E

import (
	_ "crypto/md5"
	"github.com/gogf/gf/util/guuid"
	"github.com/teris-io/shortid"
)

//生成uuid
func E取uuid() string {
	return guuid.New().String()
}

//生成短id
func E取短id() string {
	s, _ := shortid.Generate()
	return s
}
