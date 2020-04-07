package E

import (
	"fmt"
	E "github.com/duolabmeng6/goefun/core"
	"testing"
)

func TestE创建目录多级(t *testing.T) {

}

func TestE文件枚举(t *testing.T) {
	var files []string
	_ = E文件_枚举("../", ".go", &files, false, true)
	for i, value := range files {
		fmt.Println(i, value)
	}

}

func TestE目录_枚举子目录(t *testing.T) {
	var files []string
	_ = E目录_枚举子目录("../", &files, true, true)
	for i, value := range files {
		fmt.Println(i, value)
	}
}

func TestE文件_取文件名(t *testing.T) {
	var files []string
	_ = E文件_枚举("../", ".go", &files, true, true)
	for i, value := range files {
		fmt.Println(i, value, E文件_取文件名(value))

	}

}

func Test目录(t *testing.T) {
	fmt.Println(" :", E文件_取父目录(""))
	fmt.Println(". :", E文件_取父目录("."))
	fmt.Println("/a :", E文件_取父目录("/a"))
	fmt.Println("/a/b :", E文件_取父目录("/a/b"))
	fmt.Println("/a/b/ :", E文件_取父目录("/a/b/"))

	fmt.Println("/////a, /b :", E文件_路径合并处理("a", "b", "c"))
	fmt.Println("/////a, /b :", E文件_路径合并处理("a/11////", "b", "c"))

}

func TestE文件_写出(t *testing.T) {
	E.E调试输出(E.E取运行目录())

	E文件_写出(E.E取运行目录()+"/aaa/ccc/ddd/1.txt", E.E到字节集("123"))

}

func TestE文件_追加文本(t *testing.T) {
	E.E调试输出(E.E取运行目录())

	E文件_追加文本(E.E取运行目录()+"/aaa/ccc/ddd/2.txt", "123")

}

func TestE文件_保存(t *testing.T) {
	E.E调试输出(E.E取运行目录())

	E文件_保存(E.E取运行目录()+"/aaa/ccc/ddd/2.txt", "1234")

}
