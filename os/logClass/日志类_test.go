package os

import (
	"testing"
)

func TestNew日志类(t *testing.T) {
	var log = New日志类("aaa.log", "info")
	var log2 = New日志类("bbb.log", "info")

	log.Log("aaa",
		"user", "aaa",
		"pass", "123456",
	)
	log2.Log("bbb",
		"user", "bbb",
		"pass", "123456",
	)

	log.E错误日志("error",
		"user", "aaa",
		"pass", "123456",
	)

}
