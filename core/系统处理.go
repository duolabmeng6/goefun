package core

import (
	"bytes"
	"os/exec"
	"time"
)

/*
运行
打开内存文件
取剪辑板文本
置剪辑板文本
剪辑板中可有文本
清除剪辑板
取屏幕宽度
取屏幕高度
取鼠标水平位置
取鼠标垂直位置
取颜色数
输入框
信息框
取文本注册项
取数值注册项
取字节集注册
写注册项
 删除注册项
注册项是否存在
 取默认底色
快照
 读配置项
写配置项
取配置节名
取操作系统类别
多文件对话框
*/

//调用格式： 〈无返回值〉 延时 （整数型 欲等待的时间） - 系统核心支持库->系统处理
//英文名称：sleep
//本命令暂停当前程序的运行并等待指定的时间。本命令为初级命令。
//参数<1>的名称为“欲等待的时间”，类型为“整数型（int）”。本参数指定欲暂停程序执行的时间，单位为毫秒。
//
//操作系统需求： Windows、Linux
func E延时(欲等待的时间 int64) {
	time.Sleep(time.Duration(欲等待的时间) * time.Millisecond)
}

//调用格式： 〈逻辑型〉 运行 （文本型 欲运行的命令行，逻辑型 是否等待程序运行完毕，［整数型 被运行程序窗口显示方式］） - 系统核心支持库->系统处理
//英文名称：run
//本命令运行指定的可执行文件或者外部命令。如果成功，返回真，否则返回假。本命令为初级命令。
//参数<1>的名称为“欲运行的命令行”，类型为“文本型（text）”。
//参数<2>的名称为“是否等待程序运行完毕”，类型为“逻辑型（bool）”，初始值为“假”。
//参数<3>的名称为“被运行程序窗口显示方式”，类型为“整数型（int）”，可以被省略。参数值可以为以下常量之一：1、#隐藏窗口； 2、#普通激活； 3、#最小化激活； 4、#最大化激活； 5、#普通不激活； 6、#最小化不激活。如果省略本参数，默认为“普通激活”方式。
//
//操作系统需求： Windows、Linux

func E运行_win(欲运行的命令行 string, 是否等待程序运行完毕 bool) string {
	var err error
	//cmd := exec.Command("cmd")
	cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	go func(欲运行的命令行 string) {
		// start stop restart
		in.WriteString(欲运行的命令行) //写入你的命令，可以有多行，"\n"表示回车
	}(欲运行的命令行)
	err = cmd.Start()

	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}

	rt := E文本编码转换(out.String(), "gbk", "utf-8")
	//fmt.Println(rt)

	return rt
}
