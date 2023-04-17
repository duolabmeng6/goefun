package E

import (
	"github.com/duolabmeng6/efun/src/utils"
	. "github.com/duolabmeng6/goefun/ecore"
	"gopkg.in/ini.v1"
)

type E配置项 struct {
	cfg     *ini.File
	节名区分大小写 bool
}

func New配置项(配置项数据 string, 是否区分大小写 bool) *E配置项 {
	o := new(E配置项)

	if 是否区分大小写 == false {
		o.E设置节名不区分大小写()
	}

	o.E加载配置项从内存(配置项数据)
	return o
}

func (this *E配置项) E设置节名不区分大小写() {
	this.节名区分大小写 = true
}

//InsensitiveLoad
func (this *E配置项) E加载配置项从文件(fileapth string) bool {
	var err error

	this.cfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:      true,
		SpaceBeforeInlineComment: true,
		Insensitive:              this.节名区分大小写,
	}, fileapth)

	if err != nil {
		E调试输出("配置项加载失败: ", err)
		return false
	}
	return true
}
func (this *E配置项) E加载配置项从内存(v string) bool {
	var err error

	this.cfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:      true,
		SpaceBeforeInlineComment: true,
		Insensitive:              this.节名区分大小写,
	}, utils.E到字节集(v))

	if err != nil {
		E调试输出("配置项加载失败:", err)
		return false
	}
	return true
}

func (this *E配置项) E读配置项(节名称 string, 配置项名称 string, 默认文本 string) string {
	v := this.cfg.Section(节名称).Key(配置项名称).String()
	if v == "" {
		return 默认文本
	}
	return v
}

func (this *E配置项) E写配置项(节名称 string, 配置项名称 string, 设置文本 string) bool {
	this.cfg.Section(节名称).Key(配置项名称).SetValue(设置文本)
	return true
}

func (this *E配置项) E写到文件(fileapth string) {
	this.cfg.SaveTo(fileapth)
}

func (this *E配置项) E取节名() []string {
	names := this.cfg.SectionStrings()
	names = append(names[:0], names[1:]...)
	return names
}

func (this *E配置项) E取项名(节名称 string) []string {
	names := this.cfg.Section(节名称).KeyStrings()
	return names
}
