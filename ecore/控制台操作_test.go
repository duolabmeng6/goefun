// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

package ecore

import "fmt"

// ExampleE标准输出 演示E标准输出函数的使用方法
func ExampleE标准输出() {
    E标准输出("Hello world!")
    //  Output: [Hello world!]
}

// ExampleE标准输入 演示E标准输入函数的使用方法
func ExampleE标准输入() {
    fmt.Println("Please  enter  your  name:")
    name := E标准输入()
    fmt.Printf("Hello,  %s!\n", name)
}
