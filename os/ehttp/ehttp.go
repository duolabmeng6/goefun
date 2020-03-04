package ehttp

import (
	"bytes"
	"crypto/tls"
	"errors"
	. "github.com/duolabmeng6/goefun/core"
	. "github.com/duolabmeng6/goefun/coreUtil"
	"github.com/duolabmeng6/goefun/src/cookiejar"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Ehttp struct {
	client         *http.Client
	transport      *http.Transport
	Headers        http.Header
	Timeout        int
	Response       *http.Response
	Cookies        *cookiejar.Jar
	状态码            int
	cookieFilePath string
	//E重定向方式 重定向 1，不允许重定向。2，自动重定向
	E重定向方式 int
	//重定向地址
	Location string
	//E代理方式  0 使用全局代理ip访问 1 不使用代理ip访问
	E代理方式 int
	//代理ip
	Proxy string
	//全局头信息
	全局头信息 string
	//默认头信息
	默认头信息 string
	//超时时间
	TimeOut int
}

func (this *Ehttp) E设置本次不使用代理() *Ehttp {
	this.E代理方式 = 1
	return this
}
func (this *Ehttp) E设置超时时间(超时时间 int) *Ehttp {
	this.TimeOut = 超时时间
	return this
}
func (this *Ehttp) reset() *Ehttp {
	this.TimeOut = 15
	this.E重定向方式 = 2
	this.E代理方式 = 0
	return this
}

func (this *Ehttp) Get(url string, v ...interface{}) (string, bool) {
	var 附加头信息 string
	if len(v) > 1 {
		附加头信息 = E到文本(v[0])
	}

	body, _ := this.E访问(
		url,
		"GET",
		"",
		附加头信息,
	)

	return string(body), this.E访问失败()
}

//token=token&name=1.txt&file=@file:文件的绝对路径
func (this *Ehttp) Post(url string, s string, v ...interface{}) (string, bool) {
	var 附加头信息 string
	if len(v) > 1 {
		附加头信息 = E到文本(v[0])
	}

	body, _ := this.E访问(
		url,
		"POST",
		s,
		附加头信息,
	)

	return string(body), this.E访问失败()
}

func (this *Ehttp) E访问(url string, 访问方法 string, 发送文本 string, 附加头信息 string) ([]byte, error) {
	this.setObj()

	t := New时间统计类()

	var req *http.Request
	var err error
	var 文件上传头信息 string
	文件上传头信息 = ""

	if 访问方法 == "GET" {
		req, err = http.NewRequest("GET", url, nil)
	} else {
		if strings.Contains(发送文本, "@file:") {
			// File uploading request.
			buffer := new(bytes.Buffer)
			writer := multipart.NewWriter(buffer)
			for _, item := range E分割文本(发送文本, "&") {
				array := E分割文本(item, "=")
				if len(array[1]) > 6 && array[1][0:6] == "@file:" {
					path := array[1][6:]
					if !E文件是否存在(path) {
						E调试输出格式化("%s %s error:%s Time:%s ms \n", 访问方法, url, E格式化文本(`"%s" 没有找到文件`, path), t.E取毫秒())

						return nil, errors.New(E格式化文本(`"%s" 没有找到文件`, path))
					}
					name := E文件_取文件名(path)
					if file, err := writer.CreateFormFile(array[0], name); err == nil {
						if f, err := os.Open(path); err == nil {
							defer f.Close()
							if _, err = io.Copy(file, f); err != nil {
								return nil, err
							}
						} else {
							return nil, err
						}
					} else {
						return nil, err
					}
				} else {
					writer.WriteField(array[0], array[1])
				}
			}
			writer.Close()
			if req, err = http.NewRequest("POST", url, buffer); err != nil {
				return nil, err
			} else {
				文件上传头信息 = "Content-Type:" + writer.FormDataContentType()
			}
		} else {
			req, err = http.NewRequest("POST", url, strings.NewReader(发送文本))
		}

	}
	if err != nil {
		return []byte{}, err
	}

	填充来源地址 := "Referer:" + url

	_整理头信息 := this.默认头信息 + "\r\n" +
		填充来源地址 + "\r\n" +
		this.全局头信息 + "\r\n" +
		附加头信息 +
		文件上传头信息

	arr := E分割文本(_整理头信息, "\n")
	for _, v := range arr {
		kk := E删首尾空(StrCut(v, "$:"))
		vv := E删首尾空(StrCut(v, ":$"))
		if kk == "" {
			continue
		}
		req.Header.Set(kk, vv)
	}
	//让程序自动处理gzip
	req.Header.Del("Accept-Encoding")

	//client := &http.Client{}
	resp, err := this.client.Do(req)
	if err != nil {
		E调试输出格式化("%s %s error:%s Time:%s ms \n", 访问方法, url, err.Error(), t.E取毫秒())

		return []byte{}, err
	}

	this.状态码 = resp.StatusCode
	//E调试输出(this)
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	//E调试输出(E到文本(content))

	this.Response = resp

	//重定向的操作
	if this.E重定向方式 != 0 && (resp.StatusCode == 302 || resp.StatusCode == 301) {
		this.Location = resp.Header.Get("Location")
		E调试输出格式化("%s %s StatusCode:%d Time:%s ms \nLocation: %s\n", 访问方法, url, resp.StatusCode, t.E取毫秒(), this.Location)
		//自动处理重定向消息
		if this.E重定向方式 == 2 {
			return this.E访问(this.Location, "GET", 发送文本, 附加头信息)
		}
	} else {
		E调试输出格式化("%s %s StatusCode:%d Time:%s ms \n", 访问方法, url, resp.StatusCode, t.E取毫秒())
	}

	this.cookie_save()
	this.reset()
	return content, err
}

//访问失败 返回真 成功 返回假
func (this *Ehttp) E访问失败() bool {
	if this.状态码 == 200 {
		return false
	} else {
		if this.状态码 == 302 {
			return false
		}
	}
	return true

}

func (this *Ehttp) E取状态码() int {
	return this.状态码
}
func (this *Ehttp) GetLcation() string {
	return this.Response.Header.Get("Location")
}
func (this *Ehttp) E取头信息(s string) string {
	return this.Response.Header.Get(s)
}
func (this *Ehttp) E取所有头信息() string {
	str := ""
	for k, v := range this.Response.Header {
		//E调试输出P(k,v)
		str = str + E格式化文本("%v: %v\r\n", k, v[0])
	}
	return str
}
func NewHttp() *Ehttp {
	ehttp := new(Ehttp)
	ehttp.client = &http.Client{}
	ehttp.transport = &http.Transport{}
	ehttp.Cookies, _ = cookiejar.New(nil)
	ehttp.默认头信息 = `
		Accept : */*
		Accept-Language: zh-cn,
		User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36,
		Content-Type: application/x-www-form-urlencoded,
`
	return ehttp
}

func (this *Ehttp) setObj() *Ehttp {

	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(this.TimeOut) * time.Second,
			KeepAlive: time.Duration(this.TimeOut) * time.Second,
		}).DialContext,
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     2,
		IdleConnTimeout:     90 * time.Second,
		//TLSHandshakeTimeout指定等待TLS握手完成的最长时间。零值表示不设置超时。
		TLSHandshakeTimeout:   time.Duration(this.TimeOut) * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ExpectContinueTimeout: 1 * time.Second,
		//ResponseHeaderTimeout指定在发送完请求（包括其可能的主体）之后，
		//等待接收服务端的回复的头域的最大时间。零值表示不设置超时。
		ResponseHeaderTimeout: time.Duration(this.TimeOut) * time.Second,
		Proxy:                 nil,
		//如果DisableKeepAlives为真，会禁止不同HTTP请求之间TCP连接的重用。
		DisableKeepAlives:  true,
		DisableCompression: false,
	}
	//this.Proxy = "http://127.0.0.1:8888"

	if this.E代理方式 == 1 || this.Proxy == "" {
		trans.Proxy = nil
	} else {
		trans.Proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(this.Proxy)
		}
	}

	client := &http.Client{
		Transport: trans,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     this.Cookies,
		Timeout: time.Duration(this.TimeOut) * time.Second,
	}
	this.client = client
	this.transport = trans

	return this
}

// SetProxy 设置代理访问
// 	SetProxy("http://127.0.0.1:8888")
func (this *Ehttp) SetProxy(proxy string) *Ehttp {
	this.Proxy = proxy
	return this
}
func (this *Ehttp) E设置全局HTTP代理(proxy string) *Ehttp {
	return this.SetProxy(proxy)
}
func (this *Ehttp) SetTimeOut超时时间(超时时间 int) *Ehttp {
	this.TimeOut = 超时时间
	return this
}

func (this *Ehttp) E取伪造代理ip头信息() string {
	return "X-Forwarded-For: " + E取随机ip()
}
func (this *Ehttp) SetGlobalHeaders(str string) *Ehttp {
	this.全局头信息 = str
	return this
}
func (this *Ehttp) E设置全局头信息(s string) *Ehttp {
	return this.SetGlobalHeaders(s)
}
func (this *Ehttp) E设置自动管理cookie(cookie文件路径 string) *Ehttp {
	return this.SetAutoSaveCookie(cookie文件路径)
}

//设置自动保存cookie文件
func (this *Ehttp) SetAutoSaveCookie(filepath string) *Ehttp {
	this.cookieFilePath = filepath
	this.cookie_load()
	return this
}

// cookie_load 从文件中加载cookie
func (this *Ehttp) cookie_load() *Ehttp {
	if this.cookieFilePath == "" {
		return this
	}
	cdata := E读入文件(this.cookieFilePath)
	this.Cookies.JsonDeserialize(cdata)
	return this
}

// cookie_save 保存cookie到文件
func (this *Ehttp) cookie_save() *Ehttp {
	if this.cookieFilePath == "" {
		return this
	}
	cdata, _ := this.Cookies.JsonSerialize()
	if len(cdata) == 2 {
		return this
	}
	E写到文件(this.cookieFilePath, cdata)
	return this
}
