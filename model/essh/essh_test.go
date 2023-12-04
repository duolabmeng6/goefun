package ssh

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/model/eval"
	"testing"
)

func TestNewESSH(t *testing.T) {
	ecore.E加载环境变量_从文件(ecore.E取运行源文件路径() + "/.env")
	println(ecore.E读环境变量("ip"))

	ssh := NewESSH()
	err := ssh.E连接服务器(
		ecore.E读环境变量("ip"),
		eval.E到整数(ecore.E读环境变量("port")),
		ecore.E读环境变量("username"),
		ecore.E读环境变量("password"),
	)
	if err != nil {
		t.Error(err)
		return
	}
	println("连接成功")

	//执行命令, _ := ssh.E执行命令("ls -l")
	//println("ls -l", 执行命令)
	//执行命令, _ = ssh.E执行命令("pwd")
	//println("pwd", 执行命令)
	//
	//count := 0
	//err = ssh.E执行命令2("ping 127.0.0.1", func(回调信息 string) bool {
	//	count++
	//	println("ping 127.0.0.1", 回调信息)
	//	if count >= 3 {
	//		return false
	//	}
	//	return true
	//})
	//if err != nil {
	//	println(err.Error())
	//}

	//err = ssh.E执行命令2("cd nchat;docker-compose logs", func(回调信息 string) bool {
	//	println("行行行", 回调信息)
	//	return true
	//})
	//if err != nil {
	//	println(err.Error())
	//}

	//err = ssh.E上传文件(ecore.E取运行源文件路径()+"/ssh.go", "/root/abc/ssh.go", func(进度 int) {
	//	println("E上传文件进度", 进度)
	//})
	//if err != nil {
	//	println(err.Error())
	//}
	//
	//err = ssh.E下载文件("/root/abc/ssh.go", ecore.E取运行源文件路径()+"/test.txt", func(进度 int) {
	//	println("E下载文件进度", 进度)
	//})
	//if err != nil {
	//	println(err.Error())
	//}

	文件列表, err := ssh.E取文件列表("/root/abc")
	if err != nil {
		return
	}
	for _, 文件 := range 文件列表 {
		println(文件)
	}

	println("关闭")
	ssh.E关闭()
}
