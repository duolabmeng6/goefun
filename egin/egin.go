// gin 的封装
package egin

import (
	"context"
	"errors"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/egin/errorx"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gogf/gf/v2/util/gvalid"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func iauto(c *gin.Context, 规则 string, 默认值 string) string {
	//如果没有则为自动模式 从path,get,post,cookie,header,json中获取
	var 变量值 string
	if 变量值 = I(c, "path."+规则, ""); 变量值 != "" {
		return 变量值
	}
	if 变量值 = I(c, "get."+规则, ""); 变量值 != "" {
		return 变量值
	}
	if 变量值 = I(c, "post."+规则, ""); 变量值 != "" {
		return 变量值
	}
	if 变量值 = I(c, "cookie."+规则, ""); 变量值 != "" {
		return 变量值
	}
	if 变量值 = I(c, "header."+规则, ""); 变量值 != "" {
		return 变量值
	}
	//检查 content-type 如果是 application/json 则从json中获取
	if strings.Contains(c.GetHeader("Content-Type"), "application/json") {
		if 变量值 = I(c, "json."+规则, ""); 变量值 != "" {
			return 变量值
		}
	}
	return 默认值
}
func ItoInt64(c *gin.Context, 规则 string, 默认值 string) int64 {
	return ecore.E到整数(I(c, 规则, 默认值))
}

func IVerify(c *gin.Context, 规则 string, 默认值 string, 验证规则 string, 验证规则消息 string) (string, error) {
	Val := I(c, 规则, 默认值)
	err := gvalid.New().Rules(验证规则).
		Messages(验证规则消息).
		Data(Val).Run(context.Background())
	return Val, err
}

func I(c *gin.Context, 规则 string, 默认值 string) string {
	//	c.JSON(http.StatusOK, gin.H{
	//		"getid":    egin.I(c, "get.getid", "0"),
	//		"postid":   egin.I(c, "post.postid", "0"),
	//		"pathid":   egin.I(c, "path.pathid", "3"),
	//		"headerid": egin.I(c, "header.headerid", "4"),
	//		"cookieid": egin.I(c, "cookie.cookieid", "5"),
	//		"jsonid":   egin.I(c, "json.jsonid", "5"),
	//		"任意位置":     egin.I(c, "username", "5"),
	//		"password": egin.I(c, "password", "aaaaaaa"),
	//		"ccc":      egin.I(c, "cc", "cc"),
	//	})
	//分析rule规则,获取get,post,path,cookie,header,json '变量类型.变量名/修饰符',['默认值']
	//I('get.id'); 分割文本 以.分割 获取第一个元素
	var 变量值 string
	分割参数 := strings.Split(规则, ".")
	if len(分割参数) < 2 {
		return iauto(c, 规则, 默认值)
	}
	变量类型 := 分割参数[0]
	变量名 := 分割参数[1]
	//fmt.Print(变量类型, " ", 变量名, "\n")
	if 变量类型 == "get" {
		变量值 = c.Query(变量名)
	}
	if 变量类型 == "post" {
		变量值 = c.PostForm(变量名)
	}
	if 变量类型 == "path" {
		变量值 = c.Param(变量名)
	}
	if 变量类型 == "header" {
		变量值 = c.GetHeader(变量名)
	}
	if 变量类型 == "cookie" {
		变量值, _ = c.Cookie(变量名)
	}
	if 变量类型 == "json" {
		var jsonData map[string]interface{}
		if err := c.ShouldBindBodyWith(&jsonData, binding.JSON); err != nil {
			return ""
		}
		if _, ok := jsonData[变量名]; ok {
			变量值 = fmt.Sprintf("%v", jsonData[变量名])
		}
	}
	if 变量值 == "" {
		return 默认值
	} else {
		return 变量值
	}
}

// 在gin中获取所有的参数 从json,post,get中获取 优先级 json>post>get 返回map[string]interface{}
func IAll(c *gin.Context) map[string]interface{} {
	result := make(map[string]interface{})

	// 从 JSON 中获取参数
	if err := c.ShouldBindJSON(&result); err != nil {
		// 处理错误情况
	}

	// 从 POST 表单中获取参数
	if err := c.ShouldBind(&result); err != nil {
		// 处理错误情况
	}

	// 从 GET 请求中获取参数
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}

	return result
}

func E获取文件(c *gin.Context, 表单名称 string, 上传路径 string) (string, error) {
	if 上传路径 == "" {
		上传路径 = "./upload/"
	}
	file, err := c.FormFile(表单名称)
	if err != nil {
		return "", err
	}
	上传路径 = 上传路径 + file.Filename
	err = c.SaveUploadedFile(file, 上传路径)
	if err != nil {
		return "", err
	}
	return 上传路径, nil
}

func E获取多个文件(c *gin.Context, 表单名称 string, 上传路径 string) ([]string, []error) {
	if 上传路径 == "" {
		上传路径 = "./upload/"
	}
	form, _ := c.MultipartForm()
	files := form.File[表单名称+"[]"]
	var 上传文件路径 []string
	var 上传文件路径错误 []error
	for _, file := range files {
		新上传路径 := 上传路径 + file.Filename
		err := c.SaveUploadedFile(file, 新上传路径)
		上传文件路径 = append(上传文件路径, 新上传路径)
		上传文件路径错误 = append(上传文件路径错误, err)
	}
	//检查 上传文件路径错误 是否有错误 如果有则返回 如果没有错误值返回 nil
	for _, err := range 上传文件路径错误 {
		if err != nil {
			return 上传文件路径, 上传文件路径错误
		}
	}
	return 上传文件路径, nil
}

func E加载html模板路径(r *gin.Engine, 模板路径 string) {
	//加载模板路径
	//     例如 模板文件在 /view/index.html
	//     egin.E加载html模板路径(r, "view/*")
	//     使用
	//		c.HTML(http.StatusOK, "index.html", gin.H{
	//			"title": "Main website",
	//		})
	r.LoadHTMLGlob(模板路径)
}

func Verify(c *gin.Context, dst interface{}) error {
	typ := reflect.TypeOf(dst)
	// 首先判断传入参数的类型
	if !(typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct) {
		return errors.New("应使用指针传入变量 例如 验证(&user)")
	}

	// 拿到指针所指向的元素的类型
	typ = typ.Elem()
	// 拿到指针所指向的元素的值
	value := reflect.ValueOf(dst).Elem()

	// 遍历每一个字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 忽略非导出字段
		if !field.IsExported() {
			//log.Printf("field %s is not exported, ignore", field.Name)
			continue
		}

		// 判断是否设置了这个tag
		TagI := field.Tag.Get("i")
		if TagI == "" {
			continue
		}
		TagDafault := field.Tag.Get("default")
		// 如果 TagDafault 为空则使用默认值
		if TagDafault == "" {
			TagDafault = ""
		}
		TagRule := field.Tag.Get("rule")
		if TagRule == "" {
			TagRule = ""
		}
		TagMsg := field.Tag.Get("msg")
		if TagMsg == "" {
			TagMsg = ""
		}
		v, err := IVerify(c, TagI, TagDafault, TagRule, TagMsg)
		if err != nil {
			return err //遇到错误直接返回
		}
		//v := "123"

		// 根据类型来设置值
		switch fieldType := field.Type.Kind(); fieldType {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			typedV, _ := strconv.ParseInt(v, 10, 64)
			value.Field(i).SetInt(typedV)
		case reflect.String:
			value.Field(i).SetString(v)
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			typedV, _ := strconv.ParseUint(v, 10, 64)
			value.Field(i).SetUint(typedV)
		case reflect.Bool:
			value.Field(i).SetBool(v == "true")
		default:
			log.Printf("field type %s not support yet", fieldType)
		}
	}
	return nil
}

func AutoVerifyHandler(handlerFunc interface{}) func(*gin.Context) {
	// 获取参数数量
	paramNum := reflect.TypeOf(handlerFunc).NumIn()
	valueFunc := reflect.ValueOf(handlerFunc)
	funcType := reflect.TypeOf(handlerFunc)
	if funcType.Kind() != reflect.Func {
		panic("路由处理函数应为函数")
	}
	if paramNum > 1 {
		//判断第二个参数是否为指针
		if funcType.In(1).Kind() != reflect.Ptr {
			//输出当前函数的名称
			funName := runtime.FuncForPC(reflect.ValueOf(handlerFunc).Pointer()).Name()
			panic(funName + "错误控制器的第二个参数req应为指针类型 例如 Index(c *gin.Context, req *Requests.IndexRequest)")
		}
	}
	return func(context *gin.Context) {
		// 只有一个参数说明是未重构的 HandlerFunc
		if paramNum == 1 {
			valueFunc.Call(valOf(context))
		} else {
			proxyHandlerFunc(context, handlerFunc)
		}
	}
}

func proxyHandlerFunc(ctx *gin.Context, handlerFunc interface{}) {
	funcType := reflect.TypeOf(handlerFunc)
	// 获取第二个参数的类型
	typeParam := funcType.In(1).Elem()
	// 创建实例
	param := reflect.New(typeParam).Interface()
	// 绑定参数到 struct
	err := Verify(ctx, param)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status": 1,
			"code":   400,
			"msg":    err.Error(),
		})
		return
	}

	// 调用真实 HandlerFunc
	vals := reflect.ValueOf(handlerFunc).Call(valOf(ctx, param))
	// 获取返回值
	if len(vals) > 0 {
		// 获取第一个返回值
		val := vals[0]
		err := vals[1]
		// 判断返回值是否为 error
		if !err.IsNil() {
			// 判断是否为 CodeError
			switch e := err.Interface().(type) {
			case *errorx.CodeError:
				ctx.JSON(200, gin.H{
					"status": 1,
					"code":   e.Data().Code,
					"msg":    e.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"code":   400,
				"status": 1,
				"msg":    err.Interface().(error).Error(),
			})
			return
		}
		// 判断是否为 nil 如果是则不输出
		if val.IsValid() && val.IsNil() == false {
			ctx.JSON(200, val.Interface().(gin.H))
		}
	}

}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

func NewError(code int, msg string) error {
	return errorx.NewCodeError(code, msg)
}

func CopyStruct(src interface{}, dst interface{}) error {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType.Kind() != reflect.Ptr || dstType.Kind() != reflect.Ptr {
		return errors.New("参数必须为指针")
	}
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()
	if srcType.Elem().Kind() != reflect.Struct || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("参数必须为结构体指针")
	}
	for i := 0; i < srcType.Elem().NumField(); i++ {
		srcField := srcType.Elem().Field(i)
		dstField := dstType.Elem().Field(i)
		if srcField.Name != dstField.Name {
			return errors.New("结构体字段不匹配")
		}
		if srcField.Type != dstField.Type {
			return errors.New("结构体字段类型不匹配")
		}
		if !srcField.Type.AssignableTo(dstField.Type) {
			return errors.New("结构体字段类型不匹配")
		}
		dstValue.Field(i).Set(srcValue.Field(i))
	}
	return nil
}

// E绑定静态文件目录 绑定静态文件目录
// 静态目录路径 静态文件目录路径
// 例如 E绑定静态文件目录(r, "./public")
func E绑定静态文件目录(r *gin.Engine, 静态目录路径 string) {
	var 目录名数组 []string
	var 文件名数组 []string
	// 检索目录
	ecore.E目录枚举子目录(静态目录路径, &目录名数组, false, false)
	//ecore.E调试输出("目录名数组", 目录名数组)
	for _, 目录名 := range 目录名数组 {
		r.Static("/"+目录名, 静态目录路径+"/"+目录名)
	}
	//静态文件
	ecore.E文件枚举(静态目录路径, "", &文件名数组, false, false)
	//ecore.E调试输出("文件名数组", 文件名数组)
	for _, 文件名 := range 文件名数组 {
		r.StaticFile("/"+文件名, 静态目录路径+"/"+文件名)
	}

	//r.Static("/sdk", "./public/sdk")
	//r.Static("/pages", "./public/pages")
	//r.Static("/js", "./public/js")
	//r.StaticFile("/logo.png", "./public/logo.png")

}
