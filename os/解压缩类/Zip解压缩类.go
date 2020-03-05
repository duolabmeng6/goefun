package os

import (
	"github.com/gogf/gf/encoding/gcompress"
	"io"
)

//ZipPath使用压缩算法将<paths>压缩到<dest>。
//不必要的参数<prefix>表示zip文件的路径前缀。
//注意，参数<paths>可以是目录或文件，其中
//支持用“，”连接多个路径。
func Zip压缩到文件(paths, dest string, prefix ...string) error {
	return gcompress.ZipPath(paths, dest, prefix...)
}

//ZipPathWriter使用压缩算法将<paths>压缩到<writer>。
//不必要的参数<prefix>表示zip文件的路径前缀。
//注意，参数<paths>可以是目录或文件，其中
//支持用“，”连接多个路径。
func Zip压缩到io(paths string, writer io.Writer, prefix ...string) error {
	return gcompress.ZipPathWriter(paths, writer, prefix...)
}

//UnZipFile使用压缩算法将<archive>解压缩到<dest>。
//参数<path>指定<archive>的解压缩路径，
//它可用于指定要解压缩的存档文件的一部分。
//注意，参数<dest>应该是一个目录。
func Zip解压从文件(archive, dest string, path ...string) error {
	return gcompress.UnZipFile(archive, dest, path...)
}

//UnZipContent使用压缩算法将<data>解压缩到<dest>。
//参数<path>指定<archive>的解压缩路径，
//它可用于指定要解压缩的存档文件的一部分。
//注意，参数<dest>应该是一个目录。
func Zip解压从数据(data []byte, dest string, path ...string) error {
	return gcompress.UnZipContent(data, dest, path...)
}
