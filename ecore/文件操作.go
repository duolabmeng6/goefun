// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

package ecore

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// E读入文件 从指定文件名读取数据并返回读取到的字节切片。
// 如果读取失败，则返回一个空切片和错误信息。
func E读入文件(文件名 string) []byte {
	var data []byte
	data, _ = os.ReadFile(文件名)
	return data
}

// E读入文本 从指定文件名读取数据并返回读取到的文本
// 如果读取失败，则返回一个空切片和错误信息。
func E读入文本(文件名 string) string {
	var data []byte
	data, _ = os.ReadFile(文件名)
	return string(data)
}

// E写到文件 将指定的数据写入指定的文件中。如果写入成功，返回 nil。否则，返回错误信息。
func E写到文件(文件名 string, 欲写入文件的数据 []byte) error {
	父目录 := E文件取父目录(文件名)
	if !E文件是否存在(父目录) {
		E创建目录多级(父目录)
	}
	return os.WriteFile(文件名, 欲写入文件的数据, os.ModePerm)
}

// E取当前目录 返回当前程序运行的目录。
func E取当前目录() string {
	dir, err := os.Getwd()
	if err != nil {
		//提示错误 函数名 文件名 错误信息
		log.Printf("func:%s file:%s err:%s", "E取当前目录", "文件操作.go", err)
		return ""
	}
	return dir
}

// E置当前目录 设置当前程序运行的目录。
func E置当前目录(目录 string) error {
	return os.Chdir(目录)
}

// E创建目录 创建一个目录。如果创建成功，则返回 nil。否则，返回错误信息。
func E创建目录(欲创建的目录名称 string) error {
	return os.Mkdir(欲创建的目录名称, os.ModePerm)
}

// E删除目录 删除指定目录名称
//
// 参数：
// 欲删除的目录名称 string - 目录的路径名称
//
// 返回值：
// error - 操作失败会返回一个非-nil的错误对象
func E删除目录(欲删除的目录名称 string) error {
	return os.RemoveAll(欲删除的目录名称)
}

// E复制文件 复制一个文件到另一个文件
//
// 参数：
// 被复制的文件名 string - 要复制的文件路径名称
// 复制到的文件名 string - 新文件路径名称
//
// 返回值：
// error - 操作失败会返回一个非-nil的错误对象
func E复制文件(被复制的文件名 string, 复制到的文件名 string) error {
	src, err := os.Open(被复制的文件名)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(复制到的文件名, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

// E移动文件 将文件从一个位置移到另一个位置
//
// 被移动的文件 string 要被移动的文件的路径和文件名
// 移动到的位置 string 文件的目标路径和文件名
//
// 返回错误信息 error 如果发生错误会返回错误信息 成功返回nil
func E移动文件(被移动的文件 string, 移动到的位置 string) error {
	return os.Rename(被移动的文件, 移动到的位置)
}

// E删除文件 删除指定的文件
// 欲删除的文件名 string 要删除的文件的路径和文件名
// 返回错误信息 error 如果发生错误会返回错误信息 成功返回nil
func E删除文件(欲删除的文件名 string) error {
	return os.Remove(欲删除的文件名)
}

// E文件更名 重命名文件或目录
// 欲更名的原文件或目录名 string 文件的当前路径和文件名
// 欲更改为的现文件或目录名 string 文件的新路径和文件名
// 返回错误信息 error 如果发生错误会返回错误信息 成功返回nil
func E文件更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) error {
	return os.Rename(欲更名的原文件或目录名, 欲更改为的现文件或目录名)
}

// E文件是否存在 检查文件是否存在
// 欲测试的文件名称 string 要检查的文件的路径和文件名
// 返回bool值 存在返回true 不存在返回false
func E文件是否存在(欲测试的文件名称 string) bool {
	if stat, err := os.Stat(欲测试的文件名称); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// E取文件尺寸 获取文件的大小(字节)
// 文件名 string 要获取大小的文件的路径和文件名
// 返回文件大小 int64 成功返回文件大小 失败返回-1
func E取文件尺寸(文件名 string) int64 {
	f, err := os.Stat(文件名)
	if err == nil {
		return f.Size()
	} else {
		return -1
	}
}

// E取临时文件名 在指定目录中创建一个临时文件并返回其文件对象、完整路径及错误信息
// 目录名 string 要在哪个目录下创建临时文件
// 返回 *os.File 临时文件的文件对象
// 返回 filepath 临时文件的完整路径
// 返回 error 错误信息 成功时为nil
func E取临时文件名(目录名 string) (f *os.File, filepath string, err error) {
	prefix := ""
	f, err = ioutil.TempFile(目录名, prefix)
	filepath = 目录名 + f.Name()
	return f, filepath, err
}

// E复制目录 复制一个目录到另一个目录
//
//	被复制的目录名称 string - 要复制的目录路径名称
//	复制到的目录名称 string - 新目录路径名称
//
//	返回值： error - 操作失败会返回一个非-nil的错误对象
func E复制目录(被复制的目录名称 string, 复制到的目录名称 string) error {
	// 创建目标目录
	if err := os.MkdirAll(复制到的目录名称, 0755); err != nil {
		return err
	}
	// 遍历源目录
	err := filepath.Walk(被复制的目录名称, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 获取源文件/目录相对路径
		relPath, err := filepath.Rel(被复制的目录名称, path)
		if err != nil {
			return err
		}
		// 获取目标路径
		dstPath := filepath.Join(复制到的目录名称, relPath)
		// 如果是目录，创建目标目录
		if file.IsDir() {
			if err := os.MkdirAll(dstPath, file.Mode()); err != nil {
				return err
			}
			return nil
		}
		// 复制文件
		return E复制文件(path, dstPath)
	})
	return err
}

func E路径合并(elem ...string) string {
	return path.Join(elem...)
}
func E创建目录多级(欲创建的目录名称 string) error {
	return os.MkdirAll(欲创建的目录名称, os.ModePerm)
}

// E文件枚举
// 参数<1>的名称为“欲寻找的目录”，类型为“文本型”。注明：文件目录。
//
// 参数<2>的名称为“欲寻找的文件名”，类型为“文本型”。注明：如果寻找全部文件可以填入空白，.txt|.jpg找txt和jpg的文件
//
// 参数<3>的名称为“文件数组”，类型为“文本型”，接收参数数据时采用参考传递方式，允许接收空参数数据，需要接收数组数据。注明：用于装载文件数组的变量；把寻找到的文件都放在这个数组里，并返回；。
//
// 参数<4>的名称为“是否带路径”，类型为“逻辑型”，允许接收空参数数据。注明：默认为假； 真=带目录路径，如C:\012.txt； 假=不带，如 012.txt；。
//
// 参数<6>的名称为“是否遍历子目录”，类型为“逻辑型”，允许接收空参数数据。注明：留空默认为假；为真时文件数组不主动清空。
func E文件枚举(欲寻找的目录 string, 欲寻找的文件名 string, files *[]string, 是否带路径 bool, 是否遍历子目录 bool) error {
	var ok bool
	欲寻找的文件名arr := strings.Split(欲寻找的文件名, "|")
	l, err := os.ReadDir(欲寻找的目录)
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

// E目录枚举子目录
// 取一个文件夹下级子目录；成功返回子目录数量，失败返回0；通过是否枚举子目录参数，可以枚举所有的子目录
//
// 返回值类型：整数型
//
// 参数<1>的名称为“父文件夹路径”，类型为“文本型”。注明：如：D:\Program Files；目录分割符请用\，路径不以\结尾会自动添加。
//
// 参数<2>的名称为“子目录数组”，类型为“文本型”，接收参数数据时采用参考传递方式，允许接收空参数数据，需要接收数组数据。注明：用来装载返回的子目录路径；。
//
// 参数<3>的名称为“是否带路径”，类型为“逻辑型”，允许接收空参数数据。注明：可为空默认为真,假=不带，真=带;。
//
// 参数<4>的名称为“是否继续向下枚举”，类型为“逻辑型”，允许接收空参数数据。注明：为空，默认不枚举。
func E目录枚举子目录(父文件夹路径 string, 子目录数组 *[]string, 是否带路径 bool, 是否继续向下枚举 bool) error {
	l, err := os.ReadDir(父文件夹路径)
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
				err = E目录枚举子目录(tmp, 子目录数组, 是否带路径, 是否继续向下枚举)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

func E文件取文件名(路径 string, 是否需要后缀 bool) string {
	if 是否需要后缀 {
		return filepath.Base(路径)
	}
	return strings.TrimSuffix(filepath.Base(路径), filepath.Ext(路径))
}
func E文件路径合并处理(elem ...string) string {
	return path.Join(elem...)
}

func E文件取父目录(dirpath string) string {
	return path.Dir(dirpath)
}

func E文件取扩展名(filepath string) string {
	return path.Ext(filepath)
}

func E文件删除(欲删除的文件名 string) bool {
	return E删除文件(欲删除的文件名) == nil
}

// 路径不存在时自动创建
func E文件写出(文件名 string, 欲写入文件的数据 interface{}) error {
	fpath := E文件取父目录(文件名)
	if !E文件是否存在(fpath) {
		E创建目录多级(fpath)
	}
	return ioutil.WriteFile(文件名, E到字节集(欲写入文件的数据), os.ModePerm)
}

// 路径不存在时自动创建
func E文件追加文本(文件名 string, 欲追加的文本 string) error {
	fpath := E文件取父目录(文件名)
	if !E文件是否存在(fpath) {
		E创建目录多级(fpath)
	}
	file, err := os.OpenFile(文件名, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer file.Close()

	_, err = file.Write(E到字节集(欲追加的文本 + "\r\n"))
	return err
}

// 自动检查内容是否一致是否需要写出
func E文件保存(文件名 string, 欲写入文件的数据 interface{}) error {
	if E文件是否存在(文件名) {
		data := E读入文件(文件名)
		wdata := E到字节集(欲写入文件的数据)
		if !bytes.Equal(data, wdata) {
			//E调试输出("不相同写出")
			return E文件写出(文件名, wdata)
		}
		//E调试输出("内容一样不写出")
	} else {
		//E调试输出("文件不存在写出")
		return E文件写出(文件名, 欲写入文件的数据)
	}
	return nil
}
