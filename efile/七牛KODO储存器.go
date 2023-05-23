package efile

import (
	"bytes"
	"context"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type 七牛KODO储存器 struct {
	互斥锁         sync.Mutex
	七牛KODO储存器配置 七牛KODO储存器配置
}

type 七牛KODO储存器配置 struct {
	Bucket          string `json:"bucket"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	PathPrefix      string `json:"path_prefix"` //保存路径的前缀 例如 upload/
	domain          string `json:"domain"`      //域名
	Private         bool   `json:"private"`     // 是否为私有空间 真为私有空间 假为公开空间
	StorageConfig   *storage.Config
}

func New七牛KODO储存器(七牛KODO储存器配置 七牛KODO储存器配置) *七牛KODO储存器 {

	//检查路径后缀是否有 / 如果没有则加上
	if ecore.E取文本右边(七牛KODO储存器配置.PathPrefix, 1) != "/" {
		七牛KODO储存器配置.PathPrefix = 七牛KODO储存器配置.PathPrefix + "/"
	}

	return &七牛KODO储存器{
		互斥锁:         sync.Mutex{},
		七牛KODO储存器配置: 七牛KODO储存器配置,
	}
}
func (fc *七牛KODO储存器) getToken(key string) string {
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", fc.七牛KODO储存器配置.Bucket, key),
	}
	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// Put
func (fc *七牛KODO储存器) Put(key string, data interface{}) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	key = fc.七牛KODO储存器配置.PathPrefix + key
	数据 := ecore.E到字节集(data)

	formUploader := storage.NewFormUploaderEx(fc.七牛KODO储存器配置.StorageConfig, nil)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, fc.getToken(key), key, bytes.NewReader(数据), int64(len(数据)), &putExtra)

	return err
}

// Get
func (fc *七牛KODO储存器) Get(key string) ([]byte, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	key = fc.七牛KODO储存器配置.PathPrefix + key
	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)

	domain := fc.七牛KODO储存器配置.domain
	Url := ""
	if fc.七牛KODO储存器配置.Private {
		// 私有空间访问
		deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
		Url = storage.MakePrivateURL(mac, domain, key, deadline)
	} else {
		Url = storage.MakePublicURL(domain, key)
	}
	// 下载Url的内容
	resp, err := http.Get(Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	数据, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return 数据, nil
}

// Move
func (fc *七牛KODO储存器) Move(srcKey string, destKey string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	srcKey = fc.七牛KODO储存器配置.PathPrefix + srcKey
	destKey = fc.七牛KODO储存器配置.PathPrefix + destKey

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	err := bucketManager.Move(fc.七牛KODO储存器配置.Bucket, srcKey, fc.七牛KODO储存器配置.Bucket, destKey, true)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Copy
func (fc *七牛KODO储存器) Copy(srcKey string, destKey string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	srcKey = fc.七牛KODO储存器配置.PathPrefix + srcKey
	destKey = fc.七牛KODO储存器配置.PathPrefix + destKey

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	err := bucketManager.Copy(fc.七牛KODO储存器配置.Bucket, srcKey, fc.七牛KODO储存器配置.Bucket, destKey, true)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Exists
func (fc *七牛KODO储存器) Exists(key string) (bool, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	key = fc.七牛KODO储存器配置.PathPrefix + key

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	_, err := bucketManager.Stat(fc.七牛KODO储存器配置.Bucket, key)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Size
func (fc *七牛KODO储存器) Size(key string) (int64, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	key = fc.七牛KODO储存器配置.PathPrefix + key

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	stat, err := bucketManager.Stat(fc.七牛KODO储存器配置.Bucket, key)
	if err != nil {
		return 0, err
	}

	return stat.Fsize, nil
}

// List
func (fc *七牛KODO储存器) List(prefix string) ([]string, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()
	keys := []string{}

	prefix = fc.七牛KODO储存器配置.PathPrefix + prefix

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	entries, _, _, _, err := bucketManager.ListFiles(fc.七牛KODO储存器配置.Bucket, prefix, "", "", 1000)
	if err != nil {
		return keys, err
	}
	for _, entry := range entries {
		entry.Key = strings.Replace(entry.Key, fc.七牛KODO储存器配置.PathPrefix, "", 1)

		keys = append(keys, entry.Key)
	}

	return keys, nil

}

// Delete
func (fc *七牛KODO储存器) Delete(key string) error {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	key = fc.七牛KODO储存器配置.PathPrefix + key

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	err := bucketManager.Delete(fc.七牛KODO储存器配置.Bucket, key)
	if err != nil {
		return err
	}

	return nil
}

// MimeType
func (fc *七牛KODO储存器) MimeType(key string) (string, error) {
	fc.互斥锁.Lock()
	defer fc.互斥锁.Unlock()

	key = fc.七牛KODO储存器配置.PathPrefix + key

	mac := auth.New(fc.七牛KODO储存器配置.AccessKeyID, fc.七牛KODO储存器配置.AccessKeySecret)
	bucketManager := storage.NewBucketManager(mac, fc.七牛KODO储存器配置.StorageConfig)

	stat, err := bucketManager.Stat(fc.七牛KODO储存器配置.Bucket, key)
	if err != nil {
		return "", err
	}

	return stat.MimeType, nil
}
