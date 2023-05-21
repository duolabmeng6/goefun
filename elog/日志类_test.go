package elog

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
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

	fmt.Print(ecore.E读入文本("aaa.log"))
	fmt.Print(ecore.E读入文本("bbb.log"))
	ecore.E删除文件("aaa.log")
	ecore.E删除文件("bbb.log")

}
