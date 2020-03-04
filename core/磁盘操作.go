package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//取磁盘总空间
//取磁盘剩余空间
//取磁盘卷标
//置磁盘卷标
//改变驱动器
//改变目录
//取当前目录-
//创建目录-
//删除目录-
//复制文件-
//移动文件-
//删除文件-
//文件更名-
//文件是否存在-
//寻找文件
//取文件时间
//取文件尺寸
//取文件属性
//置文件属性
//取临时文件名-
//读入文件-
//写到文件-

func E取当前目录() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

//调用格式： 〈逻辑型〉 创建目录 （文本型 欲创建的目录名称） - 系统核心支持库->磁盘操作
//英文名称：MkDir
//创建一个新的目录。成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“欲创建的目录名称”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E创建目录(欲创建的目录名称 string) error {
	return os.Mkdir(欲创建的目录名称, os.ModePerm)
}

//调用格式： 〈逻辑型〉 删除目录 （文本型 欲删除的目录名称） - 系统核心支持库->磁盘操作
//英文名称：RmDir
//删除一个存在的目录及其中的所有子目录和下属文件，请务必谨慎使用本命令。成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“欲删除的目录名称”，类型为“文本型（text）”。该目录应实际存在，如果目录中存在文件或子目录，将被一并删除，因此使用本命令请千万慎重。
//
//操作系统需求： Windows、Linux

func E删除目录(欲删除的目录名称 string) error {
	return os.RemoveAll(欲删除的目录名称)
}

//调用格式： 〈逻辑型〉 复制文件 （文本型 被复制的文件名，文本型 复制到的文件名） - 系统核心支持库->磁盘操作
//英文名称：FileCopy
//成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“被复制的文件名”，类型为“文本型（text）”。
//参数<2>的名称为“复制到的文件名”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux
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

//调用格式： 〈逻辑型〉 移动文件 （文本型 被移动的文件，文本型 移动到的位置） - 系统核心支持库->磁盘操作
//英文名称：FileMove
//将文件从一个位置移动到另外一个位置。成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“被移动的文件”，类型为“文本型（text）”。
//参数<2>的名称为“移动到的位置”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux
func E移动文件(被移动的文件 string, 移动到的位置 string) error {
	return os.Rename(被移动的文件, 移动到的位置)
}

//调用格式： 〈逻辑型〉 删除文件 （文本型 欲删除的文件名） - 系统核心支持库->磁盘操作
//英文名称：kill
//成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“欲删除的文件名”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E删除文件(欲删除的文件名 string) error {
	return os.Remove(欲删除的文件名)
}

//调用格式： 〈逻辑型〉 文件更名 （文本型 欲更名的原文件或目录名，文本型 欲更改为的现文件或目录名） - 系统核心支持库->磁盘操作
//英文名称：name
//重新命名一个文件或目录。成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“欲更名的原文件或目录名”，类型为“文本型（text）”。
//参数<2>的名称为“欲更改为的现文件或目录名”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux
func E文件更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) error {
	return os.Rename(欲更名的原文件或目录名, 欲更改为的现文件或目录名)
}

//
//调用格式： 〈逻辑型〉 文件是否存在 （文本型 欲测试的文件名称） - 系统核心支持库->磁盘操作
//英文名称：IsFileExist
//判断指定的磁盘文件是否真实存在。如存在返回真，否则返回假。本命令为初级命令。
//参数<1>的名称为“欲测试的文件名称”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E文件是否存在(欲测试的文件名称 string) bool {
	if stat, err := os.Stat(欲测试的文件名称); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

//调用格式： 〈整数型〉 取文件尺寸 （文本型 文件名） - 系统核心支持库->磁盘操作
//英文名称：FileLen
//返回一个文件的长度，单位是字节。如果该文件不存在，将返回 -1。本命令为初级命令。
//参数<1>的名称为“文件名”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E取文件尺寸(文件名 string) int64 {
	f, err := os.Stat(文件名)
	if err == nil {
		return f.Size()
	} else {
		return -1
	}
}

//调用格式： 〈文本型〉 取临时文件名 （［文本型 目录名］） - 系统核心支持库->磁盘操作
//英文名称：GetTempFileName
//返回一个在指定目录中确定不存在的 .TMP 全路径文件名称。本命令为初级命令。
//参数<1>的名称为“目录名”，类型为“文本型（text）”，可以被省略。如果省略本参数，默认将使用系统的标准临时目录。
//
//操作系统需求： Windows
func E取临时文件名(目录名 string) (f *os.File, filepath string, err error) {
	prefix := ""
	f, err = ioutil.TempFile(目录名, prefix)
	filepath = 目录名 + f.Name()
	return f, filepath, err
}

//调用格式： 〈字节集〉 读入文件 （文本型 文件名） - 系统核心支持库->磁盘操作
//英文名称：ReadFile
//返回一个字节集，其中包含指定文件的所有数据。本命令为初级命令。
//参数<1>的名称为“文件名”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux
func E读入文件(文件名 string) []byte {
	var data []byte
	data, _ = ioutil.ReadFile(文件名)
	return data
}

//调用格式： 〈逻辑型〉 写到文件 （文本型 文件名，字节集 欲写入文件的数据，... ） - 系统核心支持库->磁盘操作
//英文名称：WriteFile
//本命令用作将一个或数个字节集顺序写到指定文件中，文件原有内容被覆盖。成功返回真，失败返回假。本命令为初级命令。命令参数表中最后一个参数可以被重复添加。
//参数<1>的名称为“文件名”，类型为“文本型（text）”。
//参数<2>的名称为“欲写入文件的数据”，类型为“字节集（bin）”。
//
//操作系统需求： Windows、Linux
func E写到文件(文件名 string, 欲写入文件的数据 []byte) error {
	return ioutil.WriteFile(文件名, 欲写入文件的数据, os.ModePerm)
}
