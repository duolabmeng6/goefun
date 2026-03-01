// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

package ecore

import (
	"fmt"
	"testing"
)

// TestE运行_mac 测试在Mac/Linux上运行命令
func TestE运行_mac(t *testing.T) {
	command := "ls -al"
	wait := true
	output := E运行_mac(command, wait, func(line string) {
		t.Log("实时输出:", line)
	})
	t.Log("命令输出:", output)
}

// TestE运行_win 测试在Windows上运行命令（默认禁用）
func TestE运行_win(t *testing.T) {
	return // 在非Windows系统上禁用此测试
	command := "ping"
	wait := true
	output := E运行_win(command, wait)
	fmt.Println(output)
}
