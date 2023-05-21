package etool

import (
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
)

func TestZip压缩到文件(t *testing.T) {
	ecore.E调试输出(ecore.E取当前目录())
	//Zip压缩到文件(E取当前目录()+"\\","./a.zip","efun")
	//路径前缀可以改变压缩包的文件名

	//Zip解压从文件("./a.zip","./test/","这个参数要对应上之前的前缀")
	//Zip解压从文件("./a.zip","./test/","efun")
}
