package ejs

import (
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
)

func TestRunJs(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := RunJs("function get(p){return p}", "get", "aaaaaaa"+ecore.E到文本(i))
		ecore.E调试输出(i, str)
	}

}
