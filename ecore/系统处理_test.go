package ecore

import (
	"fmt"
	"testing"
)

func TestE运行_mac(t *testing.T) {
	//command := "ping www.baidu.com"
	command := "ls -al"
	wait := true
	output := E运行_mac(command, wait, func(line string) {
		t.Log("??x", line)
	})
	t.Log("??", output)
	//  add  more  tests  here
}

func TestE运行_win(t *testing.T) {
	return
	command := "ping"
	wait := true
	output := E运行_win(command, wait)
	fmt.Println(output)
	//  add  more  tests  here
}
