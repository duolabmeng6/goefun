// 存取值用的键值对象
// 等同于Java中的HashMap 或 C#中的Dictionary对象。
package os

import (
	. "github.com/duolabmeng6/goefun/eCore"
	"github.com/duolabmeng6/goefun/gabs"
)

type EJson struct {
	Json *gabs.Container
}
type H map[string]interface{}

func New存取键值表() *EJson {
	x := new(EJson).Init()
	return x
}
func NewJson() *EJson {
	x := new(EJson).Init()
	return x
}

func (this *EJson) Init() *EJson {
	this.Json = gabs.New()

	return this
}

func (this *EJson) Clear() *EJson {
	this.Json = gabs.New()
	return this
}

func (this *EJson) Del(key string) error {
	return this.Json.DeleteP(key)
}
func (this *EJson) E删除(key string) error {
	return this.Json.DeleteP(key)
}
func (this *EJson) GetString(key string) string {
	if this.Json.ExistsP(key) {
		//str := this.Json.Path(key).Data().(string)

		return E到文本(this.Json.Path(key).Data())
	}
	return ""
}
func (this *EJson) E取文本(key string) string {
	return this.GetString(key)
}

func (this *EJson) GetInt(key string) int64 {
	if this.Json.ExistsP(key) {
		value := E到整数(this.Json.Path(key).Data())
		return value
	}
	return 0
}
func (this *EJson) E取值(key string) int64 {
	return this.GetInt(key)
}

func (this *EJson) GetBool(key string) bool {
	if this.Json.ExistsP(key) {
		value := this.Json.Path(key).Data().(bool)
		return value
	}
	return false
}
func (this *EJson) E取逻辑值(key string) bool {
	return this.GetBool(key)
}

func (this *EJson) Set(key string, value interface{}) {
	_, _ = this.Json.SetP(value, key)
}

func (this *EJson) SetArray(key string, value interface{}) {
	//_, _ = this.Json.ArrayP(key)
	_ = this.Json.ArrayAppendP(value, key)
}

func (this *EJson) Data() interface{} {
	return this.Json.Data()
}

func (this *EJson) LoadFromJsonFile(filepath string) bool {
	data := E读入文件(filepath)
	//fmt.Printf(string(data))
	var err error
	this.Json, err = gabs.ParseJSON(data)
	//fmt.Print(err)
	return err != nil
}

func (this *EJson) LoadFromJsonString(data string) bool {
	var err error
	this.Json, err = gabs.ParseJSON(E到字节集(data))
	//fmt.Print(err)
	return err != nil
}

//调用格式： 〈文本型〉 对象．到JSON （逻辑型 是否修饰，［逻辑型 是否UNICODE］，［整数型 主键排序方式］） - E2EE互联网服务器套件2.2.3->存取键值表
//英文名称：ToJSON
//内容转换到JSON文本。本命令为初级对象成员命令。
//参数<1>的名称为“是否修饰”，类型为“逻辑型（bool）”，初始值为“假”。修饰后输出的内容更便于阅读。效率会稍微低一点。
//参数<2>的名称为“是否UNICODE”，类型为“逻辑型（bool）”，可以被省略。是否输出为UNICODE文本。如果为空则默认为假，如果为真，则中文等字符将转换为UNICODE并转换为以“\u”开头带四位十六进制文本的内容。
//参数<3>的名称为“主键排序方式”，类型为“整数型（int）”，可以被省略。当前键值表和所有下级键值表输出时的主键排序方式。[0] 不排序 [1] 正向排序 [2] 反向排序。如果为空则默认为0。
//
//操作系统需求： Windows

func (this *EJson) ToJson(是否修饰 bool) string {

	if 是否修饰 {
		return this.Json.StringIndent("", "  ")
	} else {
		return this.Json.String()
	}
}
func (this *EJson) E到JSON(是否修饰 bool) string {
	return this.ToJson(是否修饰)
}
func (this *EJson) GetArrayCount(s string) int {
	int, _ := this.Json.ArrayCountP(s)
	return int
}

// GetArrayAllData 获取某个数组里面的所有数据
//
//	for i, object := range v.GetArrayAllData("array.list") {
//		fmt.Printf("i:%s value:%s \n", i, object.Path("avg").Data().(string))
//		fmt.Printf("i:%s value:%s \n", i, object.Path("avg").ToString())
//		fmt.Printf("i:%s value:%s \n", i, object.Path("avg").ToInt())
//		fmt.Printf("i:%s value:%s \n", i, object.GetString("avg"))
//	}
func (this *EJson) GetArrayAllData(s string) []*gabs.Container {
	objects := this.Json.Path(s).Children()

	return objects
}
