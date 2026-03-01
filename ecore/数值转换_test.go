// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

package ecore

import "fmt"

// Example取十六进制文本 演示E取十六进制文本函数的使用方法
func Example取十六进制文本() {
	n := 123456789
	fmt.Println(E取十六进制文本(n))
	// Output: 75bcd15
}

// Example取八进制文本 演示E取八进制文本函数的使用方法
func Example取八进制文本() {
	n := 123456789
	fmt.Println(E取八进制文本(n))
	// Output: 726746425
}

// Example十六进制 演示E十六进制函数的使用方法
func Example十六进制() {
	s := "0x499602d2"
	fmt.Println(E十六进制(s))
	// Output: 1234567890
}

// Example二进制 演示E二进制函数的使用方法
func Example二进制() {
	s := "1001001100101100000001011010010"
	fmt.Println(E二进制(s))
	// Output: 1234567890
}
