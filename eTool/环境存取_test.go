package E

import (
	"testing"

	. "github.com/duolabmeng6/goefun/eCore"
)

func TestE设置命令行(t *testing.T) {
	E调试输出(E取命令行()[1])
	E调试输出(E取命令行()[2])
	E调试输出(E取命令行()[3])

	var a, b, c string
	E设置命令行("a", "a", "这是一个a参数", &a)
	E设置命令行("b", "b", "这是一个b参数", &b)
	E设置命令行("c", "c", "这是一个c参数", &c)
	E命令行解析()

	E调试输出(a, b, c)
}
