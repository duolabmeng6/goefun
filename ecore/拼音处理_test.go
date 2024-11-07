package ecore

import (
	"reflect"
	"testing"
)

// 测试 E取所有发音
func TestE取所有发音(t *testing.T) {
	处理类 := E拼音处理类{}

	测试用例 := []struct {
		名称      string
		欲取拼音的文本 string
		包含声调    bool
		期望结果    []string
	}{
		{"单字无声调", "中", false, []string{"zhong", "zhong"}},
		{"单字有声调", "中", true, []string{"zhōng", "zhòng"}},
		{"多字取首字", "中国", false, []string{"zhong", "zhong"}},
		{"空字符串", "", false, []string{}},
	}

	for _, 用例 := range 测试用例 {
		t.Run(用例.名称, func(t *testing.T) {
			结果 := 处理类.E取所有发音(用例.欲取拼音的文本, 用例.包含声调)
			if !reflect.DeepEqual(结果, 用例.期望结果) {
				t.Errorf("期望 %v, 得到 %v", 用例.期望结果, 结果)
			}
		})
	}
}

// 测试 E取拼音
func TestE取拼音(t *testing.T) {
	测试用例 := []struct {
		名称      string
		欲取拼音的文本 string
		包含声调    bool
		启用多音字   bool
		期望结果    [][]string
	}{
		{"无声调无多音字", "中国", false, false, [][]string{{"zhong"}, {"guo"}}},
		{"有声调无多音字", "中国", true, false, [][]string{{"zhōng"}, {"guó"}}},
		{"有声调有多音字", "中国", true, true, [][]string{{"zhōng", "zhòng"}, {"guó"}}},
	}

	for _, 用例 := range 测试用例 {
		t.Run(用例.名称, func(t *testing.T) {
			结果 := E取拼音(用例.欲取拼音的文本, 用例.包含声调, 用例.启用多音字)
			if !reflect.DeepEqual(结果, 用例.期望结果) {
				t.Errorf("期望 %v, 得到 %v", 用例.期望结果, 结果)
			}
		})
	}
}

// 测试 E取首拼
func TestE取首拼(t *testing.T) {
	测试用例 := []struct {
		名称      string
		欲取拼音的文本 string
		启用多音字   bool
		期望结果    []string
	}{
		{"无多音字", "中国人", false, []string{"z", "g", "r"}},
		{"有多音字", "中国人", true, []string{"z", "g", "r"}},
	}

	for _, 用例 := range 测试用例 {
		t.Run(用例.名称, func(t *testing.T) {
			结果 := E取首拼(用例.欲取拼音的文本, 用例.启用多音字)
			if !reflect.DeepEqual(结果, 用例.期望结果) {
				t.Errorf("期望 %v, 得到 %v", 用例.期望结果, 结果)
			}
		})
	}
}
