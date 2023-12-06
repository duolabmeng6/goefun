package ezlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

// Zlib压缩数据 使用zlib算法保持压缩data
func Zlib压缩数据(data []byte) ([]byte, error) {
	if data == nil || len(data) < 13 {
		return data, nil
	}
	var in bytes.Buffer
	var err error
	w := zlib.NewWriter(&in)
	if _, err = w.Write(data); err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return in.Bytes(), err
	}
	return in.Bytes(), nil
}

// Zlib解压数据 使用zlib算法解压缩<data>。
func Zlib解压数据(data []byte) ([]byte, error) {
	if data == nil || len(data) < 13 {
		return data, nil
	}

	b := bytes.NewReader(data)
	var out bytes.Buffer
	var err error
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(&out, r); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
