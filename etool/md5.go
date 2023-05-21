package etool

import (
	"crypto/md5"
	_ "crypto/md5"
	"encoding/hex"
)

// E取md5从文本
func E取md5从文本(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// E取md5从文本
func E取md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// E取数据摘要 获取数据的md5
func E取数据摘要(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
