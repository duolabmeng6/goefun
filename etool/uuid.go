package etool

import (
	_ "crypto/md5"
	"github.com/gogf/guuid"
	"github.com/teris-io/shortid"
)

// E取UUID
func E取UUID() string {
	return guuid.New().String()
}

// E取短Id
func E取短Id() string {
	s, _ := shortid.Generate()
	return s
}
