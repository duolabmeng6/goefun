package ejson

import (
	"github.com/duolabmeng6/goefun/model/eval"
	"testing"
)

func Test_newEJson(t *testing.T) {
	ej := NewEJson()
	ej.E置值("name", "xiaoming")
	ej.E置值("age", 18)
	books := []string{"book1", "book2", "book2"}
	ej.E置值("books", books)

	friend := []map[string]interface{}{
		{"name": "xiaohong", "age": 18},
		{"name": "xiaohong", "age": 18},
	}
	ej.E置值("friends", friend)
	ej.E置值("ext.n", "1")
	ej.E置值("ext.b", "2")

	v, _ := ej.E取值("age")
	println(eval.E到文本(v))
	v, _ = ej.E取值("ext.n")
	println(eval.E到文本(v))
	v, _ = ej.E取值("books")
	println(eval.E到文本(v))
	v, _ = ej.E取值("friends")
	println(eval.E到文本(v))
	v, _ = ej.E取值("friends.0.name")
	println("friends.0.name", eval.E到文本(v))
	v, _ = ej.E取值("friends.0")
	println("friends.0", eval.E到文本(v))
	v, _ = ej.E取值("friends[0].name")
	println("friends[0].name", eval.E到文本(v))
	println(ej.E导出为JSON())

	//ej.E加载(`{"age":18,"book":["book1","book2","book2"],"name":"xiaoming"}`)
	//v, err := ej.E取值("age")
	//if err != nil {
	//	println(err.Error())
	//}
	//println(eval.E到文本(v))
	//println(ej.E导出为JSON())

}
