package gfUtil

import (
	"testing"

	E "github.com/duolabmeng6/goefun/core"
)

func Test_缓存测试(t *testing.T) {
	var cache *E缓存类 = New缓存类()
	cache.E设置("aa", "bb", 0)
	cache.E设置("a2", "b2", 0)
	E.E调试输出(cache.E获取("aa"))

	E.E调试输出(cache.E大小())
	//E.E调试输出(cache.E是否存在("aa"))
	//E.E调试输出(cache.E删除("aa"))
	//cache.E关闭()
	E.E调试输出(cache.E获取所有值())
	E.E调试输出(cache.E获取所有键())

}
