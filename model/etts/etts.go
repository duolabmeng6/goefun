package etts

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/etool"
	"github.com/hajimehoshi/go-mp3"
	"github.com/longbai/edgetts"
	"os"
)

type ETTSI interface {
	E创建() *ETTS
	E文本转语音(文本 string) error
	E设置音色(人名 string)
}

type ETTS struct {
	ETTSI
	tts             *edgetts.EdgeTTS
	displayShortMap map[string]string
	当前音色            string
	存放目录            string
}

func NewETTS(存放目录 string) *ETTS {
	e := &ETTS{}
	e.E创建()
	e.displayShortMap = map[string]string{
		"晓晓(女)":    "zh-CN-XiaoxiaoNeural",
		"晓伊(女)":    "zh-CN-XiaoyiNeural",
		"云健(中年男人)": "zh-CN-YunjianNeural",
		"云希(男孩)":   "zh-CN-YunxiNeural",
		"云夏(女孩)":   "zh-CN-YunxiaNeural",
		"云扬(男青年)":  "zh-CN-YunyangNeural",
		"女粤语1":     "zh-CN-HiuGaaiNeural",
		"女粤语2":     "zh-CN-HiuMaanNeural",
		"男粤语":      "zh-CN-WanLungNeural",
		"小北东北人":    "zh-CN-XiaobeiNeural",
	}
	e.E设置音色("晓晓(女)")
	e.存放目录 = 存放目录
	if 存放目录 == "" {
		e.存放目录 = "./tmp"
	}
	if ecore.E文件是否存在(e.存放目录) == false {
		ecore.E创建目录(e.存放目录)
	}
	//检查最后一个字符是否为/
	if e.存放目录[len(e.存放目录)-1:] != "/" {
		e.存放目录 = e.存放目录 + "/"
	}

	return e
}

func (e *ETTS) E创建() error {
	e.tts = new(edgetts.EdgeTTS)
	return nil
}

func (e *ETTS) E文本转语音(文本 string) (string, error) {

	随机文件名 := e.存放目录 + etool.E取UUID() + ".mp3"
	tts := new(edgetts.EdgeTTS)
	err := edgetts.TextToMp3(tts, 文本, e.当前音色, 随机文件名)
	if err != nil {
		return "", err
	}
	return 随机文件名, nil
}
func (e *ETTS) E设置音色(人名 string) {
	e.当前音色 = e.displayShortMap[人名]
}

// E取MP3时间 读取MP3文件的持续时间（以毫秒为单位）
func E取MP3时间(文件路径 string) (int, error) {
	file, err := os.Open(文件路径)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		return 0, err
	}

	length := decoder.Length()
	duration := float64(length) / float64(decoder.SampleRate()*4) // 4 = 2 channels * 2 bytes per sample
	//转换为毫秒
	durationInt := int(duration * 1000)
	return durationInt, nil
}
