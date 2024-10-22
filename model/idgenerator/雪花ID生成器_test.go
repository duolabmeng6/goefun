package idgenerator

import (
	"testing"
)

func TestE创建(t *testing.T) {
	// 测试有效的结点ID
	算法, err := E创建(1)
	if err != nil {
		t.Errorf("E创建(1) 失败: %v", err)
	}
	if 算法 == nil {
		t.Error("E创建(1) 返回了nil对象")
	}

	// 测试无效的结点ID
	_, err = E创建(1024)
	if err == nil {
		t.Error("E创建(1024) 应该返回错误,但没有")
	}
}

func TestE取ID(t *testing.T) {
	算法, _ := E创建(1)
	id1 := 算法.E取ID()
	id2 := 算法.E取ID()

	if id1 == id2 {
		t.Error("E取ID() 生成了重复的ID")
	}
}

func TestEID到文本(t *testing.T) {
	算法, _ := E创建(1)
	id := 算法.E取ID()

	testCases := []struct {
		编码格式 EID编码格式
		期望长度 int
	}{
		{E文本, 19},
		{Base2, 64},
		{Base32, 13},
		{Base36, 12},
		{Base58, 11},
		{Base64, 11},
	}

	for _, tc := range testCases {
		文本ID := EID到文本(id, tc.编码格式)
		if len(文本ID) != tc.期望长度 {
			t.Errorf("EID到文本(id, %v) 返回的长度为 %d, 期望长度为 %d", tc.编码格式, len(文本ID), tc.期望长度)
		}
	}
}

func TestE文本到ID(t *testing.T) {
	算法, _ := E创建(1)
	原始ID := 算法.E取ID()

	testCases := []EID编码格式{E文本, Base2, Base32, Base36, Base58, Base64}

	for _, 编码格式 := range testCases {
		文本ID := EID到文本(原始ID, 编码格式)
		转换后ID, err := E文本到ID(文本ID, 编码格式)

		if err != nil {
			t.Errorf("E文本到ID(%s, %v) 失败: %v", 文本ID, 编码格式, err)
		}

		if 转换后ID != 原始ID {
			t.Errorf("E文本到ID(%s, %v) 返回 %d, 期望 %d", 文本ID, 编码格式, 转换后ID, 原始ID)
		}
	}

	// 测试无效的输入
	_, err := E文本到ID("invalid", E文本)
	if err == nil {
		t.Error("E文本到ID(\"invalid\", E文本) 应该返回错误,但没有")
	}
}
