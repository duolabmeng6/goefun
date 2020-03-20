package E

import "fmt"

func E调试输出(a ...interface{}) {
	fmt.Println(a...)
}

func E调试输出格式化(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}
