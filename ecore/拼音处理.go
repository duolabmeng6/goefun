package ecore

import "github.com/mozillazg/go-pinyin"

//取所有发音
//取发音数目
//取拼音
//取声母
//取韵母
//发音比较
//输入字比较

// E拼音处理类 提供拼音相关处理功能，如获取单字的所有发音、获取拼音的首字母等。
type E拼音处理类 struct{}

// E取所有发音 获取单个字的所有拼音发音。
// 参数:
//   - 欲取拼音的文本: 需要获取拼音的文本，如果提供的不是单个字，则会取出首字。
//   - 包含声调: 是否包含声调，类型为bool。默认值为false。
//
// 返回值:
//   - []string: 返回该字的所有拼音发音，如果提供的文本为空或为无效字符，则返回空数组。
func (E拼音处理类) E取所有发音(欲取拼音的文本 string, 包含声调 bool) []string {
	if len(欲取拼音的文本) >= 1 {
		所有拼音 := E取拼音(string([]rune(欲取拼音的文本)[0]), 包含声调, true)
		if len(所有拼音) >= 1 {
			return 所有拼音[0]
		}
	}
	return []string{}
}

// E取拼音 获取文本的拼音。
// 参数:
//   - 欲取拼音的文本: 需要获取拼音的文本内容。
//   - 包含声调: 是否包含声调，类型为bool。默认值为false。
//   - 启用多音字: 是否启用多音字模式，类型为bool。启用后如果遇到多音字会返回多个音。默认值为false。
//
// 返回值:
//   - [][]string: 返回文本的所有拼音。如果多音字启用，则返回多个拼音，否则只返回一个拼音。
func E取拼音(欲取拼音的文本 string, 包含声调 bool, 启用多音字 bool) [][]string {
	a := pinyin.NewArgs()
	if 包含声调 {
		a.Style = pinyin.Tone
	}
	a.Heteronym = 启用多音字
	return pinyin.Pinyin(欲取拼音的文本, a)
}

// E取首拼 获取文本中所有拼音的首字母。
// 参数:
//   - 欲取拼音的文本: 需要获取拼音首字母的文本内容。
//   - 启用多音字: 是否启用多音字模式，类型为bool。启用后如果遇到多音字会返回多个音。默认值为false。
//
// 返回值:
//   - []string: 返回文本的所有拼音的首字母。
func E取首拼(欲取拼音的文本 string, 启用多音字 bool) [][]string {
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	a.Heteronym = 启用多音字
	return pinyin.Pinyin(欲取拼音的文本, a)
}
