package ejson

import (
	"github.com/duolabmeng6/goefun/model/eval"
	"reflect"
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

// Test_ini 测试E加载ini配置内容方法
func Test_ini(t *testing.T) {
	iniContent := `
		[section1]
		key1 = value1
		key2 = value2

		[section2]
		key3 = value3
	`

	ej := NewEJson()
	err := ej.E加载从env配置内容(iniContent)
	if err != nil {
		t.Errorf("E加载ini配置内容失败: %v", err)
	}

	expected := map[string]interface{}{
		"section1": map[string]interface{}{"key1": "value1", "key2": "value2"},
		"section2": map[string]interface{}{"key3": "value3"},
	}
	if !reflect.DeepEqual(ej.data, expected) {
		t.Errorf("预期结果 %+v, 实际结果 %+v", expected, ej.data)
	}
	t.Log(ej.E取值("section1.key1"))
	t.Log(ej.E取值("section1.key2"))
	t.Log(ej.E取值("section2.key3"))
}

// Test_E加载从env配置内容 测试从.env格式字符串加载内容
func Test_E加载从env配置内容(t *testing.T) {
	env内容 := `# 这是注释行
KEY1=Value1
KEY2="Value2 with quotes"
KEY3='Value3 with single quotes'
# 另一个注释行
KEY4=Value4`

	ej := NewEJson()
	err := ej.E加载从env配置内容(env内容)
	if err != nil {
		t.Fatalf("E加载从env配置内容失败: %v", err)
	}

	expected := map[string]interface{}{
		"KEY1": "Value1",
		"KEY2": "Value2 with quotes",
		"KEY3": "Value3 with single quotes",
		"KEY4": "Value4",
	}

	if !reflect.DeepEqual(ej.data, expected) {
		t.Errorf("预期结果 %+v, 实际结果 %+v", expected, ej.data)
	}

	t.Log(ej.E导出为JSON())

	value, err := ej.E取值("KEY1")
	if err != nil {
		t.Errorf("E取值 KEY1 失败: %v", err)
	} else {
		t.Logf("KEY1: %v", value)
	}

	value, err = ej.E取值("KEY2")
	if err != nil {
		t.Errorf("E取值 KEY2 失败: %v", err)
	} else {
		t.Logf("KEY2: %v", value)
	}

	value, err = ej.E取值("KEY3")
	if err != nil {
		t.Errorf("E取值 KEY3 失败: %v", err)
	} else {
		t.Logf("KEY3: %v", value)
	}
	value, err = ej.E取值("KEY4")
	if err != nil {
		t.Errorf("E取值 KEY4 失败: %v", err)
	} else {
		t.Logf("KEY4: %v", value)
	}
}
