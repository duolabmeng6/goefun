package os

import (
	"github.com/gogf/gf/util/gconv"
	"hash/crc32"
)

func Crc32加密(v interface{}) uint32 {
	return crc32.ChecksumIEEE(gconv.Bytes(v))
}
