// Package qrcodeTool Description: 二维码工具类
package qrcodeTool

import (
	"bytes"
	"image/color"
	"os"

	qrgen "github.com/skip2/go-qrcode"
	"github.com/tuotoo/qrcode"
)

// E二维码纠错级别 是二维码的纠错级别类型，使用来自 qrgen 库的原生纠错级别。
type E二维码纠错级别 = qrgen.RecoveryLevel

const (
	// E低级 表示7%的错误恢复率
	E低级 E二维码纠错级别 = qrgen.Low
	// E中级 表示15%的错误恢复率
	E中级 E二维码纠错级别 = qrgen.Medium
	// E高级 表示25%的错误恢复率
	E高级 E二维码纠错级别 = qrgen.High
	// E最高级 表示30%的错误恢复率
	E最高级 E二维码纠错级别 = qrgen.Highest
)

// E二维码工具类 提供二维码生成和识别的工具方法
type E二维码工具类 struct{}

// E生成二维码 生成一个二维码并返回其字节数组
//
// 参数:
//   - 二维码内容: 需要生成的二维码内容
//   - 尺寸: 二维码的尺寸（像素）
//   - 纠错级别: 二维码的纠错级别
//
// 返回:
//   - 二维码字节数组
//   - 错误信息（如果有）
func (E二维码工具类) E生成二维码(二维码内容 string, 尺寸 int, 纠错级别 E二维码纠错级别) ([]byte, error) {
	return qrgen.Encode(二维码内容, qrgen.RecoveryLevel(纠错级别), 尺寸)
}

// E生成并保存 生成二维码并保存为文件
//
// 参数:
//   - 二维码内容: 需要生成的二维码内容
//   - 尺寸: 二维码的尺寸（像素）
//   - 保存文件名: 保存的文件路径
//   - 纠错级别: 二维码的纠错级别
//
// 返回:
//   - 错误信息（如果有）
func (E二维码工具类) E生成并保存(二维码内容 string, 尺寸 int, 保存文件名 string, 纠错级别 E二维码纠错级别) error {
	return qrgen.WriteFile(二维码内容, qrgen.RecoveryLevel(纠错级别), 尺寸, 保存文件名)
}

// E生成并保存2 生成自定义颜色的二维码并保存为文件
//
// 参数:
//   - 二维码内容: 需要生成的二维码内容
//   - 尺寸: 二维码的尺寸（像素）
//   - 保存文件名: 保存的文件路径
//   - 背景颜色: 二维码的背景颜色
//   - 前景颜色: 二维码的前景颜色
//   - 纠错级别: 二维码的纠错级别
//
// 返回:
//   - 错误信息（如果有）
func (E二维码工具类) E生成并保存2(二维码内容 string, 尺寸 int, 保存文件名 string, 背景颜色, 前景颜色 color.Color, 纠错级别 E二维码纠错级别) error {
	return qrgen.WriteColorFile(二维码内容, qrgen.RecoveryLevel(纠错级别), 尺寸, 前景颜色, 背景颜色, 保存文件名)
}

// E识别二维码 从文件中读取并识别二维码
//
// 参数:
//   - file: 文件指针，指向需要识别的二维码图片
//
// 返回:
//   - 识别的二维码内容
//   - 错误信息（如果有）
func (E二维码工具类) E识别二维码(file *os.File) (string, error) {
	qrmatrix, err := qrcode.Decode(file)
	if err != nil {
		return "", err
	}
	return qrmatrix.Content, nil
}

// E二维码类 表示二维码对象，包括其内容和属性
type E二维码类 struct {
	E内容    string        // 二维码的内容
	E纠错级别  E二维码纠错级别      // 二维码的纠错级别
	E前景颜色  color.Color   // 二维码的前景颜色
	E背景颜色  color.Color   // 二维码的背景颜色
	E禁用边框  bool          // 是否禁用二维码的边框
	qrCode *qrgen.QRCode // 内部的二维码对象
}

// E创建 创建一个二维码对象
//
// 参数:
//   - 二维码内容: 需要生成的二维码内容
//   - 级别: 二维码的纠错级别
//
// 返回:
//   - 生成的二维码对象
//   - 错误信息（如果有）
func (E二维码类) E创建(二维码内容 string, 级别 E二维码纠错级别) (*E二维码类, error) {
	qrCode, err := qrgen.New(二维码内容, qrgen.RecoveryLevel(级别))
	if err != nil {
		return nil, err
	}
	return &E二维码类{E内容: 二维码内容, E纠错级别: 级别, qrCode: qrCode}, nil
}

// E创建2 创建一个带有指定版本号的二维码对象
//
// 参数:
//   - 二维码内容: 需要生成的二维码内容
//   - 版本号: 二维码的版本号（1-40）
//   - 级别: 二维码的纠错级别
//
// 返回:
//   - 生成的二维码对象
//   - 错误信息（如果有）
func (E二维码类) E创建2(二维码内容 string, 版本号 int, 级别 E二维码纠错级别) (*E二维码类, error) {
	qrCode, err := qrgen.NewWithForcedVersion(二维码内容, 版本号, qrgen.RecoveryLevel(级别))
	if err != nil {
		return nil, err
	}
	return &E二维码类{E内容: 二维码内容, E纠错级别: 级别, qrCode: qrCode}, nil
}

// E取坐标数组 返回二维码的坐标数组
//
// 返回:
//   - 一个二维布尔数组，表示二维码的像素点，true为黑色，false为白色
func (qr *E二维码类) E取坐标数组() [][]bool {
	return qr.qrCode.Bitmap()
}

// E生成二维码 返回指定尺寸的二维码图像
//
// 参数:
//   - 尺寸: 二维码的尺寸（像素）
//
// 返回:
//   - 二维码图像的字节数组
func (qr *E二维码类) E生成二维码(尺寸 int) []byte {
	var buf bytes.Buffer
	err := qr.qrCode.Write(尺寸, &buf)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

// E写到接口 将二维码数据写出到一个接口中
//
// 参数:
//   - 尺寸: 二维码的尺寸（像素）
//   - w: 要写入的接口（例如文件指针）
//
// 返回:
//   - 错误信息（如果有）
func (qr *E二维码类) E写到接口(尺寸 int, w *os.File) error {
	return qr.qrCode.Write(尺寸, w)
}

// E写出文件 将二维码数据保存为文件
//
// 参数:
//   - 尺寸: 二维码的尺寸（像素）
//   - 文件名: 保存的文件路径
//
// 返回:
//   - 错误信息（如果有）
func (qr *E二维码类) E写出文件(尺寸 int, 文件名 string) error {
	return qr.qrCode.WriteFile(尺寸, 文件名)
}

// E取图像文本 将二维码以文本形式返回
//
// 参数:
//   - 反向颜色: 如果为true，返回黑白反转的二维码文本
//
// 返回:
//   - 二维码的文本表示形式
func (qr *E二维码类) E取图像文本(反向颜色 bool) string {
	return qr.qrCode.ToString(反向颜色)
}

// E取小图像文本 将小尺寸二维码以文本形式返回
//
// 参数:
//   - 反向颜色: 如果为true，返回黑白反转的二维码文本
//
// 返回:
//   - 小尺寸二维码的文本表示形式
func (qr *E二维码类) E取小图像文本(反向颜色 bool) string {
	return qr.qrCode.ToSmallString(反向颜色)
}
