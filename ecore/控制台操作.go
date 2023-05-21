package ecore

import (
	"bufio"
	"fmt"
	"os"
)

// E标准输出 向控制台写入数据。
//
// 参数：
//
// 欲写入的数据 ...interface{} - 欲写入的数据
func E标准输出(欲写入的数据 ...interface{}) {
	fmt.Print(欲写入的数据, "\n")
}

// E标准输入 从控制台读取数据。
//
// 返回值： string - 读取到的数据
func E标准输入() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}
