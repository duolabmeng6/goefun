package qrcodeTool

import (
	"image/color"
	"os"
	"testing"
)

// 测试生成二维码
func TestE生成二维码(t *testing.T) {
	工具类 := E二维码工具类{}
	data, err := 工具类.E生成二维码("测试内容", 256, E中级)
	if err != nil {
		t.Fatalf("生成二维码失败: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("生成二维码的数据为空")
	}
}

// 测试生成并保存二维码
func TestE生成并保存(t *testing.T) {
	工具类 := E二维码工具类{}
	err := 工具类.E生成并保存("测试内容", 256, "test_qr.png", E中级)
	if err != nil {
		t.Fatalf("生成并保存二维码失败: %v", err)
	}

	// 确保文件生成成功
	if _, err := os.Stat("test_qr.png"); os.IsNotExist(err) {
		t.Fatalf("二维码文件未生成")
	}
	// 清理测试文件
	os.Remove("test_qr.png")
}

// 测试生成自定义颜色二维码并保存
func TestE生成并保存2(t *testing.T) {
	工具类 := E二维码工具类{}
	err := 工具类.E生成并保存2("测试内容", 256, "test_color_qr.png", color.White, color.Black, E高级)
	if err != nil {
		t.Fatalf("生成自定义颜色二维码失败: %v", err)
	}

	// 确保文件生成成功
	if _, err := os.Stat("test_color_qr.png"); os.IsNotExist(err) {
		t.Fatalf("二维码文件未生成")
	}
	// 清理测试文件
	os.Remove("test_color_qr.png")
}

// 测试二维码识别
func TestE识别二维码(t *testing.T) {
	工具类 := E二维码工具类{}

	// 先生成二维码
	err := 工具类.E生成并保存("测试识别", 256, "test_recognize_qr.png", E中级)
	if err != nil {
		t.Fatalf("生成二维码失败: %v", err)
	}

	// 打开文件进行识别
	file, err := os.Open("test_recognize_qr.png")
	if err != nil {
		t.Fatalf("打开二维码文件失败: %v", err)
	}
	defer file.Close()

	// 识别二维码
	content, err := 工具类.E识别二维码(file)
	if err != nil {
		t.Fatalf("识别二维码失败: %v", err)
	}
	if content != "测试识别" {
		t.Fatalf("二维码内容识别错误: 期望 '测试识别'，但得到 '%s'", content)
	}

	// 清理测试文件
	os.Remove("test_recognize_qr.png")
}

// 测试二维码类的创建和操作
func TestE二维码类(t *testing.T) {
	二维码类, err := E二维码类{}.E创建("测试二维码类", E中级)
	if err != nil {
		t.Fatalf("创建二维码类失败: %v", err)
	}

	data := 二维码类.E生成二维码(256)
	if len(data) == 0 {
		t.Fatalf("二维码类生成的二维码数据为空")
	}

	// 测试保存文件
	err = 二维码类.E写出文件(256, "test_class_qr.png")
	if err != nil {
		t.Fatalf("二维码类保存文件失败: %v", err)
	}

	// 确保文件生成成功
	if _, err := os.Stat("test_class_qr.png"); os.IsNotExist(err) {
		t.Fatalf("二维码类文件未生成")
	}
	// 清理测试文件
	os.Remove("test_class_qr.png")
}
