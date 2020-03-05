package os

import (
	"bytes"
	"compress/gzip"
	"io"
)

// 对数据进行GZIP压缩。
// T压缩等级 [-1] 默认压缩等级 [1] 最大压缩速度 [9] 最大压缩尺寸。
// 1 to 9 which means from none to the best compression.
//
// Note that it returns error if given <level> is invalid.
func Gzip压缩数据(data []byte, level ...int) ([]byte, error) {
	var writer *gzip.Writer
	var buf bytes.Buffer
	var err error
	if len(level) > 0 {
		writer, err = gzip.NewWriterLevel(&buf, level[0])
		if err != nil {
			return nil, err
		}
	} else {
		writer = gzip.NewWriter(&buf)
	}
	if _, err = writer.Write(data); err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 对数据进行GZIP解压。
func Gzip解压数据(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(&buf, reader); err != nil {
		return nil, err
	}
	if err = reader.Close(); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
