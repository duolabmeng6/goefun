package E

import (
	"testing"

	. "github.com/duolabmeng6/goefun/eCore"
)

func TestRunJs(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := RunJs("function get(p){return p}", "get", "aaaaaaa"+E到文本(i))
		E调试输出(i, str)
	}

}
