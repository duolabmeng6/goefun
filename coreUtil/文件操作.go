package coreUtil

import (
	. "github.com/duolabmeng6/goefun/core"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func E创建目录多级(欲创建的目录名称 string) error {
	return os.MkdirAll(欲创建的目录名称, os.ModePerm)
}

//子程序名：文件_枚举
//枚举某个目录下的指定类型文件；成功返回文件数量；
//返回值类型：整数型
//参数<1>的名称为“欲寻找的目录”，类型为“文本型”。注明：文件目录。
//参数<2>的名称为“欲寻找的文件名”，类型为“文本型”。注明：如果寻找全部文件可以填入空白，.txt|.jpg找txt和jpg的文件
//参数<3>的名称为“文件数组”，类型为“文本型”，接收参数数据时采用参考传递方式，允许接收空参数数据，需要接收数组数据。注明：用于装载文件数组的变量；把寻找到的文件都放在这个数组里，并返回；。
//参数<4>的名称为“是否带路径”，类型为“逻辑型”，允许接收空参数数据。注明：默认为假； 真=带目录路径，如C:\012.txt； 假=不带，如 012.txt；。
//参数<6>的名称为“是否遍历子目录”，类型为“逻辑型”，允许接收空参数数据。注明：留空默认为假；为真时文件数组不主动清空。
func E文件枚举(欲寻找的目录 string, 欲寻找的文件名 string, files *[]string, 是否带路径 bool, 是否遍历子目录 bool) error {
	var ok bool
	欲寻找的文件名arr := strings.Split(欲寻找的文件名, "|")
	l, err := ioutil.ReadDir(欲寻找的目录)
	if err != nil {
		return err
	}

	separator := "/"

	for _, f := range l {
		tmp := string(欲寻找的目录 + separator + f.Name())

		if f.IsDir() {
			if 是否遍历子目录 {
				err = E文件枚举(tmp, 欲寻找的文件名, files, 是否带路径, 是否遍历子目录)
				if err != nil {
					return err
				}
			}
		} else {
			ok = false
			// 目标文件类型被指定
			if !isAllEmpty(欲寻找的文件名arr) {
				// 属于目标文件类型
				if isInSuffix(欲寻找的文件名arr, f.Name()) {
					ok = true

				}
			} else { // 目标文件类型为空
				ok = true
			}
			if ok {
				if 是否带路径 {
					*files = append(*files, tmp)
				} else {
					*files = append(*files, f.Name())
				}
			}
		}
	}
	return err
}

// 判断数组各元素是否是空字符串或空格
func isAllEmpty(list []string) (isEmpty bool) {

	if len(list) == 0 {
		return true
	}

	isEmpty = true
	for _, f := range list {

		if strings.TrimSpace(f) != "" {
			isEmpty = false
			break
		}
	}

	return isEmpty
}

// 判断目标字符串的末尾是否含有数组中指定的字符串
func isInSuffix(list []string, s string) (isIn bool) {

	isIn = false
	for _, f := range list {

		if strings.TrimSpace(f) != "" && strings.HasSuffix(s, f) {
			isIn = true
			break
		}
	}

	return isIn
}

//子程序名：目录_枚举子目录1
//取一个文件夹下级子目录；成功返回子目录数量，失败返回0；通过是否枚举子目录参数，可以枚举所有的子目录
//返回值类型：整数型
//参数<1>的名称为“父文件夹路径”，类型为“文本型”。注明：如：D:\Program Files；目录分割符请用\，路径不以\结尾会自动添加。
//参数<2>的名称为“子目录数组”，类型为“文本型”，接收参数数据时采用参考传递方式，允许接收空参数数据，需要接收数组数据。注明：用来装载返回的子目录路径；。
//参数<3>的名称为“是否带路径”，类型为“逻辑型”，允许接收空参数数据。注明：可为空默认为真,假=不带，真=带;。
//参数<4>的名称为“是否继续向下枚举”，类型为“逻辑型”，允许接收空参数数据。注明：为空，默认不枚举。

func E目录_枚举子目录(父文件夹路径 string, 子目录数组 *[]string, 是否带路径 bool, 是否继续向下枚举 bool) error {
	l, err := ioutil.ReadDir(父文件夹路径)
	if err != nil {
		return err
	}
	separator := "/"
	for _, f := range l {
		tmp := string(父文件夹路径 + separator + f.Name())

		if f.IsDir() {
			if 是否带路径 {
				*子目录数组 = append(*子目录数组, tmp)
			} else {
				*子目录数组 = append(*子目录数组, f.Name())
			}
			if 是否继续向下枚举 {
				err = E目录_枚举子目录(tmp, 子目录数组, 是否带路径, 是否继续向下枚举)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

func E文件_取文件名(路径 string) string {
	return filepath.Base(路径)
}

func E文件_路径合并处理(elem ...string) string {
	return path.Join(elem...)
}

func E文件_取父目录(dirpath string) string {
	return path.Dir(dirpath)
}

func E文件_取扩展名(filepath string) string {
	return path.Ext(filepath)
}

func E文件_删除(欲删除的文件名 string) bool {
	return E删除文件(欲删除的文件名) == nil
}
func E文件_更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) bool {
	return E文件更名(欲更名的原文件或目录名, 欲更改为的现文件或目录名) == nil
}
