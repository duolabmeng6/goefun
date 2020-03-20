# GoEfun

为golang提供强大且易用的中文函数库，易语言go函数库，完整的封装了易语言核心支持库所有功能，以及易语言中简单易用的函数

# 特性

- 完全兼容window linux mac运行
- 完全兼容golang运行环境
- 全中文函数名和注释
- 封装了超100+命令
- 全部命令经过测试 100% 测试 放心使用
- 所有命令均有测试用例可以查看后立即使用
- 封装了超丰富的文本处理函数,简单,易用


## 使用方法

### 引入模块

```go
package main

import (
	. "github.com/duolabmeng6/goefun/core"
	. "github.com/duolabmeng6/goefun/coreUtil"
	. "github.com/duolabmeng6/goefun/os/ehttp"
	. "github.com/duolabmeng6/goefun/os/存取键值表"
)

func main() {
	E调试输出("欢迎使用go-efun")
	http := NewHttp()
	ret, flag := http.Get("https://www.taobao.com/help/getip.php")
	if flag {
		E调试输出("访问失败", ret)
	} else {
		E调试输出(flag, ret)
	}

	ret, flag = http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=myip")
	if flag {
		E调试输出("访问失败", ret)
	} else {
		E调试输出(flag, ret)
	}

	data := New存取键值表()
	data.Set("aaa", "111")
	data.Set("bbb", "222")
	data.Set("ccc", "333")
	data.SetArray("list", "a")
	data.SetArray("list", "b")
	data.SetArray("list", "c")

	E调试输出(data.ToJson(true))

	E调试输出(E取md5从文本("1234567"))

	E调试输出(E取现行时间().E时间到文本("Y-m-d H:i:s"))


}


```

# 文档

## 核心支持库
https://godoc.org/github.com/duolabmeng6/goefun/core

## 核心支持库工具函数库
https://godoc.org/github.com/duolabmeng6/goefun/coreUtil

## 存取键值表
https://godoc.org/github.com/duolabmeng6/goefun/os/存取键值表

## ehttp
https://godoc.org/github.com/duolabmeng6/goefun/os/ehttp


# 封装了什么函数?

## 核心支持库
```
func DBCtoSBCNew(s string) string
func E倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int
func E写到文件(文件名 string, 欲写入文件的数据 []byte) error
func E写环境变量(环境变量名称 string, 欲写入内容 string) bool
func E分割文本(待分割文本 string, 用作分割的文本 string) []string
func E创建目录(欲创建的目录名称 string) error
func E删尾空(欲删除空格的文本 string) string
func E删除文件(欲删除的文件名 string) error
func E删除目录(欲删除的目录名称 string) error
func E删首尾空(内容 string) string
func E删首空(欲删除空格的文本 string) string
func E到全角(value string) string
func E到半角(value string) string
func E到大写(value string) string
func E到字节(value interface{}) byte
func E到字节集(value interface{}) []byte
func E到小写(value string) string
func E到数值(value interface{}) float64
func E到整数(value interface{}) int64
func E到文本(value interface{}) string
func E到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error
func E取临时文件名(目录名 string) (f *os.File, filepath string, err error)
func E取代码(欲取字符代码的文本 string) int
func E取命令行() []string
func E取当前目录() string
func E取整(value interface{}) int
func E取文件尺寸(文件名 string) int64
func E取文本右边(欲取其部分的文本 string, 欲取出字符的数目 int64) string
func E取文本左边(欲取其部分的文本 string, 欲取出字符的数目 int64) string
func E取文本长度(value string) int64
func E取空白文本(重复次数 int) string
func E取绝对值(双精度小数型 float64) float64
func E取运行目录() string
func E取重复文本(重复次数 int, 待重复文本 string) string
func E取随机数(欲取随机数的最小值 int, 欲取随机数的最大值 int) int
func E四舍五入(欲被四舍五入的数值 float64, 被舍入的位置 int) float64
func E复制文件(被复制的文件名 string, 复制到的文件名 string) error
func E子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string
func E字符(字节型 int8) string
func E寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int
func E延时(欲等待的时间 int)
func E文件是否存在(欲测试的文件名称 string) bool
func E文件更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) error
func E文本编码转换(src string, oldEncoder string, newEncoder string) string
func E标准输入() string
func E标准输出(a ...interface{})
func E求余弦(欲进行计算的角 float64) float64
func E求反正切(欲求其反正切值的数值 float64) float64
func E求平方根(欲求次方数值 float64) float64
func E求次方(欲求次方数值 float64, 次方数 float64) float64
func E求正切(欲进行计算的角 float64) float64
func E求正弦(欲进行计算的角 float64) float64
func E移动文件(被移动的文件 string, 移动到的位置 string) error
func E结束()
func E置随机数种子(欲置入的种子数值 int64)
func E读入文件(文件名 string) []byte
func E读环境变量(环境变量名称 string, 默认值 ...interface{}) string
func E调试输出(a ...interface{})
func E调试输出格式化(s string, a ...interface{})
func E运行_win(欲运行的命令行 string, 是否等待程序运行完毕 bool) string
func IsSBC(c rune) bool
func SBCtoDBC(s string) string
func ToASCIIDBC(r rune) rune
func ToASCIISBC(r rune) rune
func ToDBC(c rune) rune
func ToSBC(c rune) rune
type E日期时间型
func E到时间(s string) *E日期时间型
func E到时间从时间戳(s int64) *E日期时间型
func E取现行时间() *E日期时间型
func (this *E日期时间型) E到友好时间(d *E日期时间型) string
func (this *E日期时间型) E取分钟() int64
func (this *E日期时间型) E取小时() int64
func (this *E日期时间型) E取年份() int64
func (this *E日期时间型) E取微秒() int64
func (this *E日期时间型) E取日() int64
func (this *E日期时间型) E取时间戳() int64
func (this *E日期时间型) E取时间戳微秒() int64
func (this *E日期时间型) E取时间戳毫秒() int64
func (this *E日期时间型) E取时间戳纳秒() int64
func (this *E日期时间型) E取星期几() int64
func (this *E日期时间型) E取月份() int64
func (this *E日期时间型) E取月天数() int64
func (this *E日期时间型) E取毫秒() int64
func (this *E日期时间型) E取秒() int64
func (this *E日期时间型) E取纳秒() int64
func (this *E日期时间型) E增减日期(年 int, 月 int, 日 int) *E日期时间型
func (this *E日期时间型) E增减时间(时 int, 分 int, 秒 int) *E日期时间型
func (this *E日期时间型) E大于(time *E日期时间型) bool
func (this *E日期时间型) E小于(time *E日期时间型) bool
func (this *E日期时间型) E时间到文本(format string) string
func (this *E日期时间型) E等于(time *E日期时间型) bool

```
## 核心支持库工具函数库
```
func Env(环境变量名称 string, 默认值 string) string
func E创建目录多级(欲创建的目录名称 string) error
func E判断文本(内容 string, 关键字 ...interface{}) bool
func E判断文本s(内容 string, 关键字 ...interface{}) string
func E判断文本前缀(s string, 前缀 string) bool
func E判断文本后缀(s string, 后缀 string) bool
func E到Json(v interface{}) string
func E加载环境变量_从内容到map(s string) (map[string]string, error)
func E加载环境变量_从文件(filepath string) bool
func E取md5(data []byte) string
func E取md5从文本(str string) string
func E取uuid() string
func E取文本字数(value string) int
func E取随机ip() string
func E命令行解析()
func E文件_删除(欲删除的文件名 string) bool
func E文件_取扩展名(filepath string) string
func E文件_取文件名(路径 string) string
func E文件_取父目录(dirpath string) string
func E文件_更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) bool
func E文件_路径合并处理(elem ...string) string
func E文件枚举(欲寻找的目录 string, 欲寻找的文件名 string, files *[]string, 是否带路径 bool, 是否遍历子目录 bool) error
func E文本_删中间(欲处理文本 string, 起始位置 int64, 删除长度 int64) string
func E文本_删右边(欲处理文本 string, 删除长度 int64) string
func E文本_删左边(欲处理文本 string, 删除长度 int64) string
func E文本_取出文本中汉字(s string) string
func E文本_取右边(被查找的文本 string, 欲寻找的文本 string) string
func E文本_取左边(被查找的文本 string, 欲寻找的文本 string) string
func E文本_取随机姓氏() string
func E文本_取随机字母(要取的字符数 int, 字母类型 int) string
func E文本_取随机字母和数字(要取的字符数 int) string
func E文本_取随机数字(要取的字符数 int) string
func E文本_是否为大写字母(s string) bool
func E文本_是否为字母(s string) bool
func E文本_是否为小写字母(s string) bool
func E文本_是否为数字(s string) bool
func E文本_是否为汉字(s string) bool
func E文本_自动补零(s string, len int) string
func E文本_逐字分割(s string) []string
func E文本_颠倒(s string) string
func E文本_首字母改大写(s string) string
func E文本区分_只取字母(s string) string
func E文本区分_只取数字(s string) string
func E文本区分_只取汉子(s string) string
func E文本区分_只取符号(s string) string
func E格式化文本(format string, a ...interface{}) string
func E目录_枚举子目录(父文件夹路径 string, 子目录数组 *[]string, 是否带路径 bool, 是否继续向下枚举 bool) error
func E计算数学表达式(v string) string
func E设置命令行(name string, defaultvalue string, help string, value *string)
func Json解析(data string, path string) gjson.Result
func Json解析文本(data string, path string) string
func RunJs(content string, functionName string, args ...interface{}) string
func StrCut(内容 string, 表达式 string) string
type E时间统计类
func New时间统计类() *E时间统计类
func (this *E时间统计类) E取毫秒() string
func (this *E时间统计类) E取秒() string
func (this *E时间统计类) E开始()
type E正则表达式
func New正则表达式(正则表达式文本 string, 被搜索的文本 string) (*E正则表达式, bool)
func (this *E正则表达式) E创建(正则表达式文本 string, 被搜索的文本 string) bool
func (this *E正则表达式) E取匹配数量() int
func (this *E正则表达式) E取匹配文本(匹配索引 int) string
func (this *E正则表达式) E取子匹配数量() int
func (this *E正则表达式) E取子匹配文本(匹配索引 int, 子表达式索引 int) string
func (this *E正则表达式) GetResult() [][]string
type E线程池
func New线程池(最大任务数量 int) *E线程池
func (this *E线程池) E初始化(最大任务数量 int)
func (this *E线程池) E加入任务()
func (this *E线程池) E完成()
func (this *E线程池) E等待任务完成()
type E读写共享锁
func (this *E读写共享锁) E开始写()
func (this *E读写共享锁) E开始读()
func (this *E读写共享锁) E结束写()
func (this *E读写共享锁) E结束读()
type E配置项
func New配置项(配置项数据 string, 是否区分大小写 bool) *E配置项
func (this *E配置项) E写到文件(fileapth string)
func (this *E配置项) E写配置项(节名称 string, 配置项名称 string, 设置文本 string) bool
func (this *E配置项) E加载配置项从内存(v string) bool
func (this *E配置项) E加载配置项从文件(fileapth string) bool
func (this *E配置项) E取节名() []string
func (this *E配置项) E取项名(节名称 string) []string
func (this *E配置项) E设置节名不区分大小写()
func (this *E配置项) E读配置项(节名称 string, 配置项名称 string, 默认文本 string) string
type StrArray
func New文本型数组() *StrArray
func (a *StrArray) E删除成员(欲删除的位置 int)
func (a *StrArray) E加入成员(s string)
func (a *StrArray) E取(i int) string
func (a *StrArray) E取数组成员数() int
func (a *StrArray) E插入成员(欲插入的位置 int, 欲插入的成员数据 string)
func (a *StrArray) E数组排序(排序方向是否为从小到大 bool)
func (a *StrArray) E清除数组()
```


## 存取键值表

```
type EJson
func NewJson() *EJson
func New存取键值表() *EJson
func (this *EJson) Clear() *EJson
func (this *EJson) Data() interface{}
func (this *EJson) E到JSON(是否修饰 bool) string
func (this *EJson) E取值(key string) int64
func (this *EJson) E取文本(key string) string
func (this *EJson) E取逻辑值(key string) bool
func (this *EJson) GetArrayAllData(s string) []*gabs.Container
func (this *EJson) GetArrayCount(s string) int
func (this *EJson) GetBool(key string) bool
func (this *EJson) GetInt(key string) int64
func (this *EJson) GetString(key string) string
func (this *EJson) Init() *EJson
func (this *EJson) LoadFromJsonFile(filepath string) bool
func (this *EJson) LoadFromJsonString(data string) bool
func (this *EJson) Set(key string, value interface{})
func (this *EJson) SetArray(key string, value interface{})
func (this *EJson) ToJson(是否修饰 bool) string
type H
```

## ehttp

```
type Ehttp
func NewHttp() *Ehttp
func (this *Ehttp) E取伪造代理ip头信息() string
func (this *Ehttp) E取头信息(s string) string
func (this *Ehttp) E取所有头信息() string
func (this *Ehttp) E取状态码() int
func (this *Ehttp) E设置全局HTTP代理(proxy string) *Ehttp
func (this *Ehttp) E设置全局头信息(s string) *Ehttp
func (this *Ehttp) E设置本次不使用代理() *Ehttp
func (this *Ehttp) E设置自动管理cookie(cookie文件路径 string) *Ehttp
func (this *Ehttp) E设置超时时间(超时时间 int) *Ehttp
func (this *Ehttp) E访问(url string, 访问方法 string, 发送文本 string, 附加头信息 string) ([]byte, error)
func (this *Ehttp) E访问失败() bool
func (this *Ehttp) Get(url string, v ...interface{}) (string, bool)
func (this *Ehttp) GetLcation() string
func (this *Ehttp) Post(url string, s string, v ...interface{}) (string, bool)
func (this *Ehttp) SetAutoSaveCookie(filepath string) *Ehttp
func (this *Ehttp) SetGlobalHeaders(str string) *Ehttp
func (this *Ehttp) SetProxy(proxy string) *Ehttp
func (this *Ehttp) SetTimeOut超时时间(超时时间 int) *Ehttp
```