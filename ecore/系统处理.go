package ecore

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"time"
)

func E延时(欲等待的时间 int64) {
	time.Sleep(time.Duration(欲等待的时间) * time.Millisecond)
}

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
	if 是否等待程序运行完毕 {
		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	}

	rt := E文本编码转换(out.String(), "gbk", "utf-8")
	//fmt.Println(rt)

	return rt
}

// 模拟终端的输入输出
func E运行_mac(欲运行的命令行 string, 是否等待程序运行完毕 bool, fc interface{}) string {
	// 启动一个新的进程运行命令
	cmd := exec.Command("bash", "-c", 欲运行的命令行)

	// 获取命令的输出和错误
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// 实时检查输出,并调用回调函数
	output := ""
	if fc != nil {
		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				line := scanner.Text()
				fc.(func(string))(line)
				output += line + "\n"
			}
		}()
	}

	// 运行命令
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// 等待命令完成
	if 是否等待程序运行完毕 {
		err = cmd.Wait()
		if err != nil {
			// 读取并返回错误输出
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				line := scanner.Text()
				output += line + "\n"
			}
		}
	}

	return output
}
