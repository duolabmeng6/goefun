package E

import (
	"testing"

	. "github.com/duolabmeng6/goefun/eCore"
)

func TestNew正则表达式(t *testing.T) {
	str := `aaa111bbb444ccc
	aaa222bbb555ccc
	aaa333bbb666ccc`

	var zz E正则表达式
	if zz.E创建(`aaaaa(.*?)bbb(.*?)ccc`, str) {
		E调试输出(zz.E取匹配数量())
		for i := 0; i < zz.E取匹配数量(); i++ {
			E调试输出(i, zz.E取子匹配文本(i, 1))
			E调试输出(i, zz.E取子匹配文本(i, 2))
			E调试输出(zz.E取子匹配文本(i, 3))
		}
		E调试输出(zz.E取子匹配文本(4, 3))
	}

	zz2, flag := New正则表达式(`aaa(.*?)bbb(.*?)ccc`, str)
	E调试输出(flag)
	E调试输出(zz2.E取匹配数量())
	for i := 0; i < zz2.E取匹配数量(); i++ {
		E调试输出(i, zz2.E取子匹配文本(i, 1))
		E调试输出(i, zz2.E取子匹配文本(i, 2))
		E调试输出(zz2.E取子匹配文本(i, 3))
	}
	E调试输出(zz2.E取子匹配文本(4, 3))

	arr := zz2.GetResult()

	E调试输出("GetResult", arr)

	zz3, flag := New正则表达式(`aaa`, str)

	arr2 := zz3.E替换("666")

	E调试输出("E替换", arr2)

}
