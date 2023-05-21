package ecore

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func ExampleE写到文件() {
	err := E写到文件("./1.txt", []byte("123"))
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleE读入文件() {
	内容 := E读入文件("./1.txt")
	fmt.Print(内容)
}

func ExampleE删除文件() {
	err := E删除文件("./1.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestE文件读写删(t *testing.T) {
	err := E写到文件("./1.txt", []byte("123"))
	// 抛出错误
	if err != nil {
		t.Errorf("文件读写删() = %v, want %v", err, nil)
	}
	内容 := E读入文件("./1.txt")
	err = E删除文件("./1.txt")
	// 抛出错误
	if err != nil {
		t.Errorf("文件读写删() = %v, want %v", err, nil)
	}
	// 断言 期望值和实际值是否相等
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("文件读写删() = %v, want %v", 内容, []byte("123"))
	}
}
func ExampleE取当前目录() {
	当前目录 := E取当前目录()
	fmt.Print(当前目录)
}
func TestE取当前目录(t *testing.T) {
	当前目录 := E取当前目录()
	if 当前目录 == "" {
		t.Errorf("E取当前目录() = %v, want %v", 当前目录, "不为空")
	}
}

func ExampleE创建目录() {
	欲创建的目录名称 := "./test"
	err := E创建目录(欲创建的目录名称)
	if err != nil {
		log.Fatal(err)
	}
}
func TestE创建目录(t *testing.T) {
	欲创建的目录名称 := "./test"
	err := E创建目录(欲创建的目录名称)
	if err != nil {
		t.Errorf("E创建目录() = %v, want %v", err, nil)
	}
	E删除目录(欲创建的目录名称)

}

func ExampleE删除目录() {
	欲删除的目录名称 := "./test"
	err := E删除目录(欲删除的目录名称)
	if err != nil {
		log.Fatal(err)
	}
}
func TestE删除目录(t *testing.T) {
	欲删除的目录名称 := "./test"
	err := E删除目录(欲删除的目录名称)
	if err != nil {
		t.Errorf("E删除目录() = %v, want %v", err, nil)
	}
}

func ExampleE复制文件() {
	被复制的文件名 := "./1.txt"
	复制到的文件名 := "./2.txt"
	err := E复制文件(被复制的文件名, 复制到的文件名)
	if err != nil {
		log.Fatal(err)
	}
}
func TestE复制文件(t *testing.T) {
	被复制的文件名 := "./1.txt"
	复制到的文件名 := "./2.txt"
	E写到文件("./1.txt", []byte("123"))

	err := E复制文件(被复制的文件名, 复制到的文件名)
	if err != nil {
		t.Errorf("E复制文件() = %v, want %v", err, nil)
	}
	内容 := E读入文件("./2.txt")
	err = E删除文件("./1.txt")
	err = E删除文件("./2.txt")
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("E复制文件() = %v, want %v", 内容, []byte("123"))
	}
}

func ExampleE复制目录() {
	被复制的目录名 := "./test"
	复制到的目录名 := "./test2"
	err := E复制目录(被复制的目录名, 复制到的目录名)
	if err != nil {
		log.Fatal(err)
	}

}

func TestE复制目录(t *testing.T) {
	被复制的目录名 := "./test"
	复制到的目录名 := "./test2"
	E创建目录(被复制的目录名)
	E写到文件(被复制的目录名+"/1.txt", []byte("123"))

	err := E复制目录(被复制的目录名, 复制到的目录名)
	if err != nil {
		log.Fatal(err)
	}
	内容 := E读入文件(复制到的目录名 + "/1.txt")
	err = E删除目录(被复制的目录名)
	err = E删除目录(复制到的目录名)

	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("E复制文件() = %v, want %v", 内容, []byte("123"))
	}
}

func ExampleE移动文件() {
	被移动的文件名 := "./1.txt"
	移动到的文件名 := "./2.txt"
	err := E移动文件(被移动的文件名, 移动到的文件名)
	if err != nil {
		log.Fatal(err)
	}
}

func TestE移动文件(t *testing.T) {
	被移动的文件名 := "./1.txt"
	移动到的文件名 := "./2.txt"
	E写到文件("./1.txt", []byte("123"))

	err := E移动文件(被移动的文件名, 移动到的文件名)
	if err != nil {
		t.Errorf("E移动文件() = %v, want %v", err, nil)
	}
	内容 := E读入文件("./2.txt")
	err = E删除文件("./2.txt")
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("E移动文件() = %v, want %v", 内容, []byte("123"))
	}

}

func ExampleE文件更名() {
	被更名的文件名 := "./1.txt"
	更名到的文件名 := "./2.txt"
	err := E文件更名(被更名的文件名, 更名到的文件名)
	if err != nil {
		log.Fatal(err)
	}
}

func TestE文件更名(t *testing.T) {
	被更名的文件名 := "./1.txt"
	更名到的文件名 := "./2.txt"
	E写到文件("./1.txt", []byte("123"))

	err := E文件更名(被更名的文件名, 更名到的文件名)
	if err != nil {
		t.Errorf("E文件更名() = %v, want %v", err, nil)
	}
	内容 := E读入文件("./2.txt")
	err = E删除文件("./2.txt")
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("E文件更名() = %v, want %v", 内容, []byte("123"))
	}
}

// E文件是否存在
func ExampleE文件是否存在() {
	文件名 := "./1.txt"
	文件是否存在 := E文件是否存在(文件名)
	fmt.Println(文件是否存在)
	// Output:
	// false
}

func TestE文件是否存在(t *testing.T) {
	文件名 := "./1.txt"
	文件是否存在 := E文件是否存在(文件名)
	if 文件是否存在 {
		t.Errorf("E文件是否存在() = %v, want %v", 文件是否存在, false)
	}

}

// E取文件尺寸
func ExampleE取文件尺寸() {
	文件名 := "./1.txt"
	文件尺寸 := E取文件尺寸(文件名)
	fmt.Println(文件尺寸)
}

func TestE取文件尺寸(t *testing.T) {
	文件名 := "./1.txt"
	E写到文件(文件名, []byte("123"))
	文件尺寸 := E取文件尺寸(文件名)
	if 文件尺寸 != 3 {
		t.Errorf("E取文件尺寸() = %v, want %v", 文件尺寸, 0)
	}
	E删除文件(文件名)

}

// E取临时文件名
func ExampleE取临时文件名() {
	_, 临时文件名, _ := E取临时文件名("")
	fmt.Println(临时文件名)
}

func TestE取临时文件名(t *testing.T) {
	_, 临时文件名, err := E取临时文件名("")
	if err != nil {
		t.Errorf("E取临时文件名() = %v, want %v", err, false)
	}
	E删除文件(临时文件名)

}

func TestE文件枚举(t *testing.T) {
	var files []string
	_ = E文件_枚举("../", ".go", &files, false, true)
	for i, value := range files {
		fmt.Println(i, value)
	}

}

func TestE目录枚举子目录(t *testing.T) {
	var files []string
	_ = E目录枚举子目录("../", &files, true, true)
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
	E调试输出(E取运行目录())

	E文件_写出(E取运行目录()+"/aaa/ccc/ddd/1.txt", E到字节集("123"))

}

func TestE文件_追加文本(t *testing.T) {
	E调试输出(E取运行目录())

	E文件_追加文本(E取运行目录()+"/aaa/ccc/ddd/2.txt", "123")

}

func TestE文件_保存(t *testing.T) {
	E调试输出(E取运行目录())

	E文件_保存(E取运行目录()+"/aaa/ccc/ddd/2.txt", "1234")

}
