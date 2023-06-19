// Package etranslation 翻译模块
package etranslation

// 翻译模块
type 翻译模块 interface {
	E取初始化参数() []string
	E翻译(text, from, to string) (string, error)
}

// 翻译结构体
type E翻译 struct {
	服务 map[string]翻译模块
}

func New翻译() *E翻译 {
	return &E翻译{
		服务: make(map[string]翻译模块),
	}
}

func (t *E翻译) E注册服务(name string, service 翻译模块) {
	t.服务[name] = service
}

func (t *E翻译) E取翻译模块(name string) 翻译模块 {
	return t.服务[name]
}

func (t *E翻译) E模块是否存在(name string) bool {
	_, ok := t.服务[name]
	return ok
}
func (t *E翻译) E列出翻译模块() []string {
	模块列表 := make([]string, 0, len(t.服务))
	for name := range t.服务 {
		模块列表 = append(模块列表, name)
	}
	return 模块列表
}

type 翻译模块初始化参数 struct {
	Name   string
	Params []string `json:"params"`
}

func (t *E翻译) E列出翻译模块和初始化参数() []翻译模块初始化参数 {
	// 生成模块列表
	// {
	//     name:"deepL翻译",
	//     Param:["appId","appKey"],
	// },
	模块列表 := make([]翻译模块初始化参数, 0, len(t.服务))
	for name, s := range t.服务 {
		模块 := 翻译模块初始化参数{
			Name:   name,
			Params: s.E取初始化参数(),
		}

		模块列表 = append(模块列表, 模块)
	}

	return 模块列表
}
