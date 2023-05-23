package efile

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/duolabmeng6/goefun/ecore"
	"io"
	"strconv"
	"strings"
	"sync"
)

type 阿里云OSS储存器 struct {
	互斥锁         sync.Mutex
	Oss         *oss.Client
	阿里云OSS储存器配置 阿里云OSS储存器配置
}
type 阿里云OSS储存器配置 struct {
	Bucket          string `json:"bucket"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	PathPrefix      string `json:"path_prefix"` //保存路径的前缀 例如 upload/
	OssConfig       *oss.Config
}

func New阿里云OSS储存器(阿里云OSS储存器配置 阿里云OSS储存器配置) *阿里云OSS储存器 {
	client, err := oss.New(阿里云OSS储存器配置.Endpoint, 阿里云OSS储存器配置.AccessKeyID, 阿里云OSS储存器配置.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	//检查路径后缀是否有 / 如果没有则加上
	if ecore.E取文本右边(阿里云OSS储存器配置.PathPrefix, 1) != "/" {
		阿里云OSS储存器配置.PathPrefix = 阿里云OSS储存器配置.PathPrefix + "/"
	}

	return &阿里云OSS储存器{
		互斥锁:         sync.Mutex{},
		Oss:         client,
		阿里云OSS储存器配置: 阿里云OSS储存器配置,
	}
}

// Put
func (fc *阿里云OSS储存器) Put(key string, data interface{}) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return err
	}
	// 上传文件。
	数据 := ecore.E到字节集(data)
	err = bucket.PutObject(fc.阿里云OSS储存器配置.PathPrefix+key, bytes.NewReader(数据))
	if err != nil {
		return err
	}
	return nil
}

// Get
func (fc *阿里云OSS储存器) Get(key string) ([]byte, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return nil, err
	}
	// 下载文件到流。
	body, err := bucket.GetObject(fc.阿里云OSS储存器配置.PathPrefix + key)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	数据, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return 数据, err
}

// Delete
func (fc *阿里云OSS储存器) Delete(key string) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return err
	}
	// 删除文件。
	err = bucket.DeleteObject(fc.阿里云OSS储存器配置.PathPrefix + key)
	if err != nil {
		return err
	}
	return nil
}

// Move
func (fc *阿里云OSS储存器) Move(key string, tokey string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return false, err
	}
	// 移动文件。
	_, err = bucket.CopyObject(fc.阿里云OSS储存器配置.PathPrefix+key, fc.阿里云OSS储存器配置.PathPrefix+tokey)
	if err != nil {
		return false, err
	}
	// 删除文件
	err = bucket.DeleteObject(fc.阿里云OSS储存器配置.PathPrefix + key)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Copy
func (fc *阿里云OSS储存器) Copy(key string, tokey string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return false, err
	}
	// 移动文件。
	_, err = bucket.CopyObject(fc.阿里云OSS储存器配置.PathPrefix+key, fc.阿里云OSS储存器配置.PathPrefix+tokey)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Exists
func (fc *阿里云OSS储存器) Exists(key string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return false, err
	}
	// 判断文件是否存在。
	found, err := bucket.IsObjectExist(fc.阿里云OSS储存器配置.PathPrefix + key)
	if err != nil {
		return false, err
	}
	return found, nil
}

// Size
func (fc *阿里云OSS储存器) Size(key string) (int64, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return 0, err
	}
	// 获取文件基本信息。
	props, err := bucket.GetObjectMeta(fc.阿里云OSS储存器配置.PathPrefix + key)
	if err != nil {
		return 0, err
	}
	size, err := strconv.ParseInt(props.Get("content-length"), 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// List
func (fc *阿里云OSS储存器) List(key string) ([]string, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	keys := []string{}

	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return keys, err
	}
	// 获取文件列表。
	marker := oss.Marker("")
	prefix := oss.Prefix(fc.阿里云OSS储存器配置.PathPrefix + key)
	lsRes, err := bucket.ListObjectsV2(marker, prefix)
	if err != nil {
		return keys, err
	}

	for _, object := range lsRes.Objects {
		// object.Key 替换 PathPrefix 为空
		object.Key = strings.Replace(object.Key, fc.阿里云OSS储存器配置.PathPrefix, "", 1)
		keys = append(keys, object.Key)
	}
	return keys, nil
}

// MimeType
func (fc *阿里云OSS储存器) MimeType(key string) (string, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	// 创建存储空间。
	bucket, err := fc.Oss.Bucket(fc.阿里云OSS储存器配置.Bucket)
	if err != nil {
		return "", err
	}
	// 获取文件基本信息。
	props, err := bucket.GetObjectMeta(fc.阿里云OSS储存器配置.PathPrefix + key)
	if err != nil {
		return "", err
	}
	return props.Get("content-type"), nil

}
