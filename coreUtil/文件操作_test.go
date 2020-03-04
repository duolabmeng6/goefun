package coreUtil

import (
	"fmt"
	"testing"
)

func TestE创建目录多级(t *testing.T) {

}

func TestE文件枚举(t *testing.T) {
	var files []string
	_ = E文件枚举("../", ".go", &files, false, true)
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
	_ = E文件枚举("../", ".go", &files, true, true)
	for i, value := range files {
		fmt.Println(i, value, E文件_取文件名(value))

	}

}
