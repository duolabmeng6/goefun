package os

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestCrc32加密(t *testing.T) {
	gtest.Case(t, func() {
		s := "pibigstar"
		result := 693191136
		encrypt1 := Crc32加密(s)
		encrypt2 := Crc32加密([]byte(s))
		gtest.AssertEQ(int(encrypt1), result)
		gtest.AssertEQ(int(encrypt2), result)

		strmd5, _ := gmd5.Encrypt(s)
		test1 := Crc32加密(strmd5)
		test2 := Crc32加密([]byte(strmd5))
		gtest.AssertEQ(test2, test1)
	})
}
