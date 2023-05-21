package ecore

import "fmt"

func Example取十六进制文本() {
	n := 123456789
	fmt.Println(E取十六进制文本(n))
	// Output: 75bcd15
}

// 示例测试函数
func Example取八进制文本() {
	n := 123456789
	fmt.Println(E取八进制文本(n))
	// Output: 726746425
}

// 示例测试函数
func Example十六进制() {
	s := "0x499602d2"
	fmt.Println(E十六进制(s))
	// Output: 1234567890
}

// 示例测试函数
func Example二进制() {
	s := "1001001100101100000001011010010"
	fmt.Println(E二进制(s))
	// Output: 1234567890
}
