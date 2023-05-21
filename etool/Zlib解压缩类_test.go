package etool

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestZlib压缩(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := Zlib压缩数据([]byte(src))
		gtest.Assert(data, dst)

		data, _ = Zlib解压数据(dst)
		gtest.Assert(data, []byte(src))

		data, _ = Zlib压缩数据(nil)
		gtest.Assert(data, nil)
		data, _ = Zlib解压数据(nil)
		gtest.Assert(data, nil)

		data, _ = Zlib解压数据(dst[1:])
		gtest.Assert(data, nil)
	})

}
