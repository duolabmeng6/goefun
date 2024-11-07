package ebarcode

import (
	"testing"
)

/*
package main

import "awesomeProject7/ebarcode"

	func main() {
		data := []byte("Test Aztec Code")
		minError := 5
		layers := 5
		b, err := ebarcode.E生成二维条码(data, minError, layers)
		if err != nil {
			println("E生成二维条码 失败: %v", err)
		}
		ebarcode.E保存条码到文件(b, "1.png")
		a := ebarcode.E条码{E条码: b}
		a.E置尺寸(300, 300)
		a.E保存文件("2.png")
	}
*/
func TestE生成二维条码(t *testing.T) {
	data := []byte("Test Aztec Code")
	minError := 5
	layers := 5
	_, err := E生成二维条码(data, minError, layers)
	if err != nil {
		t.Errorf("E生成二维条码 失败: %v", err)
	}
	//E保存条码到文件(b, "1.png")
	//a := E条码{条码: b}
	//a.E置尺寸(300, 300)
	//a.E保存文件("2.png")
}

func TestE生成Codabar(t *testing.T) {
	content := "A123456A"
	_, err := E生成Codabar(content)
	if err != nil {
		t.Errorf("E生成Codabar 失败: %v", err)
	}
}

func TestE生成Code39(t *testing.T) {
	content := "CODE39"
	includeChecksum := true
	asciiMode := true
	_, err := E生成Code39(content, includeChecksum, asciiMode)
	if err != nil {
		t.Errorf("E生成Code39 失败: %v", err)
	}
}

func TestE生成Code93(t *testing.T) {
	content := "CODE93"
	includeChecksum := true
	asciiMode := true
	_, err := E生成Code93(content, includeChecksum, asciiMode)
	if err != nil {
		t.Errorf("E生成Code93 失败: %v", err)
	}
}

func TestE生成Code128(t *testing.T) {
	content := "CODE128"
	_, err := E生成Code128(content)
	if err != nil {
		t.Errorf("E生成Code128 失败: %v", err)
	}
}

func TestE生成Code128无校验(t *testing.T) {
	content := "CODE128-NO-CHECKSUM"
	_, err := E生成Code128无校验(content)
	if err != nil {
		t.Errorf("E生成Code128无校验 失败: %v", err)
	}
}

func TestE生成矩阵条码(t *testing.T) {
	content := "DataMatrix"
	_, err := E生成矩阵条码(content)
	if err != nil {
		t.Errorf("E生成矩阵条码 失败: %v", err)
	}
}

func TestE生成EAN(t *testing.T) {
	content := "123456789012"
	_, err := E生成EAN(content)
	if err != nil {
		t.Errorf("E生成EAN 失败: %v", err)
	}
}

func TestE生成PDF417(t *testing.T) {
	content := "PDF417"
	errorCorrectionLevel := byte(4)
	_, err := E生成PDF417(content, errorCorrectionLevel)
	if err != nil {
		t.Errorf("E生成PDF417 失败: %v", err)
	}
}

func TestE生成二五条码(t *testing.T) {
	content := "123456" // 必须是偶数位
	interleaved := true
	_, err := E生成二五条码(content, interleaved)
	if err != nil {
		t.Errorf("E生成二五条码 失败: %v", err)
	}
}
func TestE生成二五条码NonInterleaved(t *testing.T) {
	content := "12345" // 奇数位可以在非交错模式下使用
	interleaved := false
	_, err := E生成二五条码(content, interleaved)
	if err != nil {
		t.Errorf("E生成二五条码 (非交错模式) 失败: %v", err)
	}
}
