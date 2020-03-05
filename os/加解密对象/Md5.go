package os

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/gconv"
)

func Md5加密(data interface{}) (encrypt string, err error) {
	return gmd5.EncryptBytes(gconv.Bytes(data))
}

func Md5加密从文件(path string) (encrypt string, err error) {
	return gmd5.EncryptFile(path)
}
