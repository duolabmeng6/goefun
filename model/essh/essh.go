// Package ssh Description: ssh连接服务器 上传文件 下载文件 执行命令
package ssh

import (
	"bufio"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"strings"
)

type ESSHI interface {
	E连接服务器(服务器IP地址 string, 端口 int, 用户名 string, 密码 string) error
	E执行命令(命令 string) (string, error)
	E执行命令2(命令 string, fn func(回调信息 string) bool) error
	E关闭() error
	E上传文件(本地文件 string, 远程文件 string, fn func(进度 int)) error
	E下载文件(远程文件 string, 本地文件 string, fn func(进度 int)) error
	E取文件列表(远程文件夹 string) ([]string, error)
}

type ESSH struct {
	ESSHI
	client     *ssh.Client
	sftpClient *sftp.Client
}

func NewESSH() *ESSH {
	return &ESSH{
		client:     nil,
		sftpClient: nil,
	}
}

func (e *ESSH) E连接服务器(服务器IP地址 string, 端口 int, 用户名 string, 密码 string) error {
	config := &ssh.ClientConfig{
		User: 用户名,
		Auth: []ssh.AuthMethod{
			ssh.Password(密码),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var err error
	e.client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", 服务器IP地址, 端口), config)
	if err != nil {
		return err
	}
	return err
}

func (e *ESSH) E执行命令(命令 string) (string, error) {
	session, err := e.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(命令)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
func (e *ESSH) E执行命令2(命令 string, fn func(回调信息 string) bool) error {
	session, err := e.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	err = session.Start(命令)
	if err != nil {
		return err
	}

	// 创建协程来处理输出
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			if !fn(scanner.Text()) { // 如果回调函数返回 false，则停止执行
				session.Close() // 尝试关闭会话来停止命令
				return
			}
		}
	}()

	return session.Wait() // 等待命令执行完成或被中断
}

func (e *ESSH) E关闭() error {
	if e.sftpClient != nil {
		e.sftpClient.Close()
	}
	return e.client.Close()
}

func (e *ESSH) _sftpContent() error {
	if e.sftpClient == nil {
		var err error
		e.sftpClient, err = sftp.NewClient(e.client)
		if err != nil {
			return err
		}
	}
	return nil
}
func (e *ESSH) E上传文件(本地文件 string, 远程文件 string, fn func(进度 int)) error {
	err := e._sftpContent()
	if err != nil {
		return err
	}

	localFile, err := os.Open(本地文件)
	if err != nil {
		return err
	}
	defer localFile.Close()

	//检查远程文件夹是否存在
	fdir := ecore.E文件取父目录(远程文件)
	_, err = e.sftpClient.Stat(fdir)
	if err != nil {
		//文件夹不存在
		err = e.sftpClient.Mkdir(fdir)
		if err != nil {
			return err
		}
	}
	remoteFile, err := e.sftpClient.Create(远程文件)
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	stat, err := localFile.Stat()
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	bytesRead := 0
	totalBytesRead := 0
	for {
		n, err := localFile.Read(buffer)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		bytesRead += n
		_, err = remoteFile.Write(buffer[:n])
		if err != nil {
			return err
		}
		totalBytesRead += n
		progress := int(float32(totalBytesRead) / float32(stat.Size()) * 100)
		fn(progress)
	}
	return nil
}
func (e *ESSH) E下载文件(远程文件 string, 本地文件 string, fn func(进度 int)) error {
	err := e._sftpContent()
	if err != nil {
		return err
	}
	sftp, err := sftp.NewClient(e.client)
	if err != nil {
		return err
	}
	defer sftp.Close()

	remoteFile, err := sftp.Open(远程文件)
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	stat, err := remoteFile.Stat()
	if err != nil {
		return err
	}

	localFile, err := os.Create(本地文件)
	if err != nil {
		return err
	}
	defer localFile.Close()

	buffer := make([]byte, 1024)
	totalBytesRead := 0
	for {
		n, err := remoteFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		_, err = localFile.Write(buffer[:n])
		if err != nil {
			return err
		}
		totalBytesRead += n
		progress := int(float64(totalBytesRead) / float64(stat.Size()) * 100)
		fn(progress)
	}

	// 确保在文件传输完成时报告100%
	fn(100)

	return nil
}

func (e *ESSH) E取文件列表(远程文件夹 string) ([]string, error) {
	err := e._sftpContent()
	if err != nil {
		return nil, err
	}
	files, err := e.sftpClient.ReadDir(远程文件夹)
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList, nil
}
