package ecache

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
)

func Example文件缓存器() {
	// 创建缓存适配器
	缓存 := New缓存类(New文件缓存器("./cache/"))

	// 设置缓存数据
	缓存.Set("a", "1", -160)
	缓存.Set("b", 2, 60)
	缓存.Set("c", []byte("3"), 60)

	// 获取缓存数据
	a := 缓存.Get("a")
	fmt.Println(a)
	if a != nil {
		panic("a != nil")
	}
	b := 缓存.GetInt("b")
	fmt.Println(b)
	if b != 2 {
		panic("b != 2")
	}
	c := 缓存.GetBytes("c")
	fmt.Println(c)
	if string(c) != "3" {
		panic("string(c) != 3")
	}
	d := 缓存.Get("d")
	fmt.Println(d)
	if d != nil {
		panic("d != nil")
	}
	ecore.E删除目录("./cache/")
}

func Example内存缓存器() {
	// 创建缓存适配器
	缓存 := New缓存类(New内存缓存器())

	// 设置缓存数据
	缓存.Set("a", "1", -160)
	缓存.Set("b", 2, 60)
	缓存.Set("c", []byte("3"), 60)

	// 获取缓存数据
	a := 缓存.Get("a")
	fmt.Println(a)
	if a != nil {
		panic("a != nil")
	}
	b := 缓存.GetInt("b")
	fmt.Println(b)
	if b != 2 {
		panic("b != 2")
	}
	c := 缓存.GetBytes("c")
	fmt.Println(c)
	if string(c) != "3" {
		panic("string(c) != 3")
	}
	d := 缓存.Get("d")
	fmt.Println(d)
	if d != nil {
		panic("d != nil")
	}
}

func ExampleSqlite缓存器() {
	// 创建缓存适配器
	缓存 := New缓存类(NewSqlite缓存器("./cache.db"))

	// 设置缓存数据
	缓存.Set("a", "1", -160)
	缓存.Set("b", 2, 60)
	缓存.Set("c", []byte("3"), 60)

	// 获取缓存数据
	a := 缓存.Get("a")
	fmt.Println(a)
	if a != nil {
		panic("a != nil")
	}
	b := 缓存.GetInt("b")
	fmt.Println(b)
	if b != 2 {
		panic("b != 2")
	}
	c := 缓存.GetBytes("c")
	fmt.Println(c)
	if string(c) != "3" {
		panic("string(c) != 3")
	}
	d := 缓存.Get("d")
	fmt.Println(d)
	if d != nil {
		panic("d != nil")
	}
	ecore.E删除文件("./cache.db")
}

func ExampleMysql缓存器() {
	// 创建缓存适配器
	缓存 := New缓存类(NewMysql缓存器("root@tcp(127.0.0.1:3310)/gogorm?charset=utf8&parseTime=true&loc=Local"))

	// 设置缓存数据
	缓存.Set("a", "1", -160)
	缓存.Set("b", 2, 60)
	缓存.Set("c", []byte("3"), 60)

	// 获取缓存数据
	a := 缓存.Get("a")
	fmt.Println(a)
	if a != nil {
		panic("a != nil")
	}
	b := 缓存.GetInt("b")
	fmt.Println(b)
	if b != 2 {
		panic("b != 2")
	}
	c := 缓存.GetBytes("c")
	fmt.Println(c)
	if string(c) != "3" {
		panic("string(c) != 3")
	}
	d := 缓存.Get("d")
	fmt.Println(d)
	if d != nil {
		panic("d != nil")
	}
}

func ExampleRedis缓存器() {
	// 创建缓存适配器
	缓存 := New缓存类(NewRedis缓存器("127.0.0.1:6379", "", 1))

	// 设置缓存数据
	缓存.Set("a", "1", -160)
	缓存.Set("b", 2, 60)
	缓存.Set("c", []byte("3"), 60)

	// 获取缓存数据
	a := 缓存.Get("a")
	fmt.Println(a)
	if a != nil {
		panic("a != nil")
	}
	b := 缓存.GetInt("b")
	fmt.Println(b)
	if b != 2 {
		panic("b != 2")
	}
	c := 缓存.GetBytes("c")
	fmt.Println(c)
	if string(c) != "3" {
		panic("string(c) != 3")
	}
	d := 缓存.Get("d")
	fmt.Println(d)
	if d != nil {
		panic("d != nil")
	}
}
