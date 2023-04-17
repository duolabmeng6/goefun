package E

import (
	"testing"

	. "github.com/duolabmeng6/goefun/ecore"
)

func TestNew文本型数组(t *testing.T) {
	strarr := New文本型数组()
	E调试输出(strarr)
	strarr.E加入成员("3")
	strarr.E加入成员("2")
	strarr.E加入成员("5")
	strarr.E加入成员("1")
	strarr.E加入成员("4")
	strarr.E插入成员(1, "6")

	E调试输出(strarr)
	E调试输出(strarr.E取数组成员数())
	strarr.E删除成员(1)
	E调试输出(strarr.E取数组成员数())

	//strarr.E清除数组()

	strarr.E数组排序(true)
	E调试输出(strarr)
	strarr.E数组排序(false)
	E调试输出(strarr)
	E调试输出(strarr.Get(0))
	E调试输出(strarr.Array[0])

	for i := 0; i < strarr.E取数组成员数(); i++ {
		//哪一个写法好?
		E调试输出(i, strarr.Get(i))
		E调试输出(i, strarr.E取(i))
		E调试输出(i, strarr.Array[i])
	}
}
