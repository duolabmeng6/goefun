package ezip

import (
	. "github.com/duolabmeng6/goefun/ecore"
	"io"

	"github.com/gogf/gf/v2/encoding/gcompress"
)

// Zip压缩到文件 使用压缩算法将<压缩包路径>压缩到<目标目录>。
// 不必要的参数<路径前缀>表示zip文件的路径路径前缀。
// 注意，参数<压缩包路径>可以是目录或文件，其中
// 支持用“，”连接多个路径。
func Zip压缩到文件(压缩包路径, 目标目录 string, 路径前缀 ...string) error {
	writer, 临时文件, err := E取临时文件名("")
	if err != nil {
		return err
	}
	err = gcompress.ZipPathWriter(压缩包路径, writer, 路径前缀...)
	writer.Close()
	E移动文件(临时文件, 目标目录)
	return err
}

// Zip压缩到io 使用压缩算法将<压缩包路径>压缩到<writer>。
// 不必要的参数<路径前缀>表示zip文件的路径路径前缀。
// 注意，参数<压缩包路径>可以是目录或文件，其中
// 支持用“，”连接多个路径。
func Zip压缩到io(压缩包路径 string, writer io.Writer, 路径前缀 ...string) error {
	return gcompress.ZipPathWriter(压缩包路径, writer, 路径前缀...)
}

// Zip解压从文件 使用压缩算法将<压缩包路径>解压缩到<目标目录>。
// 参数<路径前缀>指定<压缩包路径>的解压缩路径，
// 它可用于指定要解压缩的存档文件的一部分。
// 注意，参数<目标目录>应该是一个目录。
func Zip解压从文件(压缩包路径, 目标目录 string, 路径前缀 ...string) error {
	return gcompress.UnZipFile(压缩包路径, 目标目录, 路径前缀...)
}

// Zip解压从数据 使用压缩算法将<data>解压缩到<目标目录>。
// 参数<路径前缀>指定<压缩包路径>的解压缩路径，
// 它可用于指定要解压缩的存档文件的一部分。
// 注意，参数<目标目录>应该是一个目录。
func Zip解压从数据(data []byte, 目标目录 string, 路径前缀 ...string) error {
	return gcompress.UnZipContent(data, 目标目录, 路径前缀...)
}
