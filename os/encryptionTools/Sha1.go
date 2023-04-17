package os

import (
	"github.com/gogf/gf/crypto/gsha1"
)

func Sha1加密(v interface{}) string {
	return gsha1.Encrypt(v)
}

func Sha1加密从文件(path string) (encrypt string, err error) {
	return gsha1.EncryptFile(path)
}
