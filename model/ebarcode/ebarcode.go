// Package ebarcode Description: 条码工具类
package ebarcode

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/aztec"
	"github.com/boombuler/barcode/codabar"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/code93"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/twooffive"
	"image/png"
	"os"
)

// E生成二维条码 生成Aztec二维条码
// 参数:
//   - 欲生成的数据: []byte 需要生成的内容数据
//   - 最小错误: int 最小的错误校正码百分比 (5%-95%)
//   - 层数: int 生成Aztec码的层数 (最多32层)
//
// 返回值:
//   - barcode.Barcode: 生成的Aztec条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成二维条码(欲生成的数据 []byte, 最小错误 int, 层数 int) (barcode.Barcode, error) {
	return aztec.Encode(欲生成的数据, 最小错误, 层数)
}

// E生成Codabar 生成Codabar条码
// 参数:
//   - 条码内容: string 生成条码的内容
//
// 返回值:
//   - barcode.Barcode: 生成的Codabar条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成Codabar(条码内容 string) (barcode.Barcode, error) {
	return codabar.Encode(条码内容)
}

// E生成Code39 生成Code39条码
// 参数:
//   - 条码内容: string 生成条码的内容
//   - 包含校验和: bool 是否包含校验和
//   - ASCII模式: bool 是否启用ASCII编码模式
//
// 返回值:
//   - barcode.Barcode: 生成的Code39条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成Code39(条码内容 string, 包含校验和 bool, ASCII模式 bool) (barcode.Barcode, error) {
	return code39.Encode(条码内容, 包含校验和, ASCII模式)
}

// E生成Code93 生成Code93条码
// 参数:
//   - 条码内容: string 生成条码的内容
//   - 包含校验和: bool 是否包含校验和
//   - ASCII模式: bool 是否启用ASCII编码模式
//
// 返回值:
//   - barcode.Barcode: 生成的Code93条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成Code93(条码内容 string, 包含校验和 bool, ASCII模式 bool) (barcode.Barcode, error) {
	return code93.Encode(条码内容, 包含校验和, ASCII模式)
}

// E生成Code128 生成带校验和的Code128条码
// 参数:
//   - 条码内容: string 生成条码的内容
//
// 返回值:
//   - barcode.Barcode: 生成的Code128条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成Code128(条码内容 string) (barcode.Barcode, error) {
	return code128.Encode(条码内容)
}

// E生成Code128无校验 生成不带校验和的Code128条码
// 参数:
//   - 条码内容: string 生成条码的内容
//
// 返回值:
//   - barcode.Barcode: 生成的无校验和的Code128条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成Code128无校验(条码内容 string) (barcode.Barcode, error) {
	return code128.EncodeWithoutChecksum(条码内容)
}

// E生成矩阵条码 生成Datamatrix矩阵条码
// 参数:
//   - 条码内容: string 生成条码的内容
//
// 返回值:
//   - barcode.Barcode: 生成的Datamatrix条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成矩阵条码(条码内容 string) (barcode.Barcode, error) {
	return datamatrix.Encode(条码内容)
}

// E生成EAN 生成EAN条码 (根据内容自动生成EAN8或EAN13)
// 参数:
//   - 条码内容: string 生成条码的内容
//
// 返回值:
//   - barcode.Barcode: 生成的EAN条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成EAN(条码内容 string) (barcode.Barcode, error) {
	return ean.Encode(条码内容)
}

// E生成PDF417 生成PDF417条码
// 参数:
//   - 条码内容: string 生成条码的内容
//   - 纠错级别: byte 错误校正级别 (范围0-8)
//
// 返回值:
//   - barcode.Barcode: 生成的PDF417条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成PDF417(条码内容 string, 纠错级别 byte) (barcode.Barcode, error) {
	return pdf417.Encode(条码内容, 纠错级别)
}

// E生成二五条码 生成2of5条码
// 参数:
//   - 条码内容: string 生成条码的内容
//   - 启用交错编码: bool 是否启用交错编码
//
// 返回值:
//   - barcode.Barcode: 生成的2of5条码对象
//   - error: 如果生成过程中发生错误，则返回错误
func E生成二五条码(条码内容 string, 启用交错编码 bool) (barcode.Barcode, error) {
	return twooffive.Encode(条码内容, 启用交错编码)
}

// E保存条码到文件 保存条码到指定文件路径
// 参数:
//   - 条码对象: barcode.Barcode 需要保存的条码对象
//   - 文件名: string 保存条码的目标文件名
//
// 返回值:
//   - error: 如果保存过程中发生错误，则返回错误
func E保存条码到文件(条码对象 barcode.Barcode, 文件名 string) error {
	// 打开文件
	file, err := os.Create(文件名)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将条码对象以PNG格式写入文件
	err = png.Encode(file, 条码对象)
	if err != nil {
		return err
	}

	return nil
}

// E条码接口 定义了通用条码的接口方法
type E条码接口 interface {
	E置尺寸(宽度, 高度 int) (barcode.Barcode, error)
	E保存文件(文件名 string) error
	E到字节数组() ([]byte, error)
	E取元数据() barcode.Metadata
	E取内容() string
}

// 实现E条码接口的结构体
type E条码 struct {
	E条码 barcode.Barcode
}

// E置尺寸 设置条码尺寸
func (e *E条码) E置尺寸(宽度, 高度 int) (barcode.Barcode, error) {
	// 调用barcode库的Scale方法调整尺寸
	scaled, err := barcode.Scale(e.E条码, 宽度, 高度)
	if err != nil {
		return nil, err
	}
	e.E条码 = scaled
	return e.E条码, nil
}

// E保存文件 保存条码到文件
func (e *E条码) E保存文件(文件名 string) error {
	return E保存条码到文件(e.E条码, 文件名)
}

// E到字节数组 将条码转换为字节数组
func (e *E条码) E到字节数组() ([]byte, error) {
	// 使用内存缓冲区保存条码图像
	var buf []byte
	buffer := bytes.NewBuffer(buf)
	err := png.Encode(buffer, e.E条码)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// E取元数据 获取条码的元数据
func (e *E条码) E取元数据() barcode.Metadata {
	return e.E条码.Metadata()
}

// E取内容 获取条码的内容
func (e *E条码) E取内容() string {
	return e.E条码.Content()
}
