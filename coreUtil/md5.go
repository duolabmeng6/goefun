package coreUtil

import (
	"crypto/md5"
	_ "crypto/md5"
	"encoding/hex"
)

func E取md5从文本(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func E取md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
