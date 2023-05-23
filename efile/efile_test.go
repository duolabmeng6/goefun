package efile

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleNew文件储存类() {
	文件系统 := New文件储存类("local")
	文件系统.设置储存器("local", New本地文件储存器("./file/"))
	文件系统.Put("a.txt", "abc")
	数据, _ := 文件系统.Get("a.txt")
	println(string(数据))
	// Output: abc
}
func ExampleNew阿里云OSS储存器() {
	ecore.E加载环境变量_从文件("./.env")
	文件系统 := New文件储存类("oss")
	文件系统.设置储存器("oss", New阿里云OSS储存器(阿里云OSS储存器配置{
		Endpoint:        "oss-cn-guangzhou.aliyuncs.com",
		Bucket:          "你的储存桶名称",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("aliyun_access_key"),
		AccessKeySecret: ecore.E读环境变量("aliyun_access_secret"),
	}))
	文件系统.Put("a.txt", "abc")
	数据, _ := 文件系统.Get("a.txt")
	println(string(数据))
	// Output: abc
}

func ExampleNew七牛云储存器() {
	ecore.E加载环境变量_从文件("./.env")
	文件系统 := New文件储存类("qiniu")
	文件系统.设置储存器("qiniu", New七牛KODO储存器(七牛KODO储存器配置{
		Bucket:          "你的储存桶名称",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("qiniu_access_key"),
		AccessKeySecret: ecore.E读环境变量("qiniu_access_secret"),
		StorageConfig: &storage.Config{
			// 空间对应的机房
			Zone: &storage.Zone_as0,
			// 是否使用https域名
			UseHTTPS: false,
			// 上传是否使用CDN上传加速
			UseCdnDomains: false,
		},
		domain:  "http://rv3k84oie.sabkt.gdipper.com",
		Private: true,
	}))
	文件系统.Put("a.txt", "abc")
	数据, _ := 文件系统.Get("a.txt")
	println(string(数据))
	// Output: abc
}

func Example文件储存类(t *testing.T) {
	ecore.E加载环境变量_从文件("./.env")

	文件系统 := New文件储存类("local")
	文件系统.设置储存器("local", New本地文件储存器("./file/"))
	文件系统.设置储存器("oss", New阿里云OSS储存器(阿里云OSS储存器配置{
		Endpoint:        "oss-cn-guangzhou.aliyuncs.com",
		Bucket:          "testupload123",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("aliyun_access_key"),
		AccessKeySecret: ecore.E读环境变量("aliyun_access_secret"),
	}))
	文件系统.设置储存器("qiniu", New七牛KODO储存器(七牛KODO储存器配置{
		Bucket:          "testgoefun",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("qiniu_access_key"),
		AccessKeySecret: ecore.E读环境变量("qiniu_access_secret"),
		StorageConfig: &storage.Config{
			// 空间对应的机房
			Zone: &storage.Zone_as0,
			// 是否使用https域名
			UseHTTPS: false,
			// 上传是否使用CDN上传加速
			UseCdnDomains: false,
		},
		domain:  "http://rv3k84oie.sabkt.gdipper.com",
		Private: true,
	}))

	// 这里可以用中文的函数
	文件系统.E储存器("local").E保存文件("a.txt", "abc")
	// 也可以用英文的函数
	文件系统.Disk("local").Put("a.txt", "abc")
	// 默认储存器是 New文件储存类("local") 里面设置的储存器
	文件系统.Put("a.txt", "abc")
	//如果想改变默认的储存器,可以这样设置
	文件系统.当前储存器 = "oss"
	//这里是oos的储存器操作的
	文件系统.Put("a.txt", "abc")

}

func Test本地文件储存器(t *testing.T) {
	文件系统 := New文件储存类("local")
	文件系统.设置储存器("local", New本地文件储存器("./file/"))

	err := 文件系统.Put("a.txt", "abc")
	assert.Equal(t, nil, err, "Put 错误")

	数据, err := 文件系统.Get("a.txt")
	assert.Equal(t, nil, err, "Get 错误")
	assert.Equal(t, "abc", string(数据), "Get 错误")

	mimetype, err := 文件系统.MimeType("a.txt")
	assert.Equal(t, nil, err, "MimeType 错误")
	assert.Equal(t, "application/octet-stream", mimetype, "MimeType 错误")

	move, err := 文件系统.Move("a.txt", "b.txt")
	assert.Equal(t, nil, err, "Move 错误")
	assert.Equal(t, true, move, "Move 错误")

	copyx, err := 文件系统.Copy("b.txt", "c.txt")
	assert.Equal(t, nil, err, "Copy 错误")
	assert.Equal(t, true, copyx, "Copy 错误")

	Exists, err := 文件系统.Exists("c.txt")
	assert.Equal(t, nil, err, "Exists 错误")
	assert.Equal(t, true, Exists, "Exists 错误")

	size, _ := 文件系统.Size("c.txt")
	assert.Equal(t, int64(3), size, "Size 错误")

	list, err := 文件系统.List("")
	assert.Equal(t, nil, err, "List 错误")
	assert.Equal(t, []string{"./file/b.txt", "./file/c.txt"}, list, "List 错误")

	err = 文件系统.Delete("a.txt")
	// 断言 err 会出现异常
	assert.NotNil(t, err, "Delete 错误")

	err = 文件系统.Delete("b.txt")
	assert.Equal(t, nil, err, "Delete 错误")
	err = 文件系统.Delete("c.txt")
	assert.Equal(t, nil, err, "Delete 错误")

	list, err = 文件系统.List("")
	assert.Equal(t, nil, err, "List 错误")
	assert.Equal(t, []string{}, list, "List 错误")
	ecore.E删除目录("./file")
}

func Test阿里云OSS储存器(t *testing.T) {
	ecore.E加载环境变量_从文件("./.env")

	文件系统 := New文件储存类("oss")
	文件系统.设置储存器("oss", New阿里云OSS储存器(阿里云OSS储存器配置{
		Endpoint:        "oss-cn-guangzhou.aliyuncs.com",
		Bucket:          "testupload123",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("aliyun_access_key"),
		AccessKeySecret: ecore.E读环境变量("aliyun_access_secret"),
	}))
	err := 文件系统.Put("a.txt", "abc")
	assert.Equal(t, nil, err, "Put 错误")

	数据, err := 文件系统.Get("a.txt")
	assert.Equal(t, nil, err, "Get 错误")
	assert.Equal(t, "abc", string(数据), "Get 错误")

	mimetype, err := 文件系统.MimeType("a.txt")
	assert.Equal(t, nil, err, "MimeType 错误")
	assert.Equal(t, "", mimetype, "MimeType 错误")

	move, err := 文件系统.Move("a.txt", "b.txt")
	assert.Equal(t, nil, err, "Move 错误")
	assert.Equal(t, true, move, "Move 错误")

	copyx, err := 文件系统.Copy("b.txt", "c.txt")
	assert.Equal(t, nil, err, "Copy 错误")
	assert.Equal(t, true, copyx, "Copy 错误")

	Exists, err := 文件系统.Exists("c.txt")
	assert.Equal(t, nil, err, "Exists 错误")
	assert.Equal(t, true, Exists, "Exists 错误")

	size, _ := 文件系统.Size("c.txt")
	assert.Equal(t, int64(3), size, "Size 错误")

	list, err := 文件系统.List("")
	assert.Equal(t, nil, err, "List 错误")
	assert.Equal(t, []string{"test/b.txt", "test/c.txt"}, list, "List 错误")

	err = 文件系统.Delete("a.txt")
	// 断言 err 会出现异常
	//assert.NotNil(t, err, "Delete 错误")

	err = 文件系统.Delete("b.txt")
	assert.Equal(t, nil, err, "Delete 错误")
	err = 文件系统.Delete("c.txt")
	assert.Equal(t, nil, err, "Delete 错误")

	list, err = 文件系统.List("")
	assert.Equal(t, nil, err, "List 错误")
	assert.Equal(t, []string{}, list, "List 错误")
}

func Test七牛KODO储存器(t *testing.T) {
	ecore.E加载环境变量_从文件("./.env")
	文件系统 := New文件储存类("qiniu")
	文件系统.设置储存器("qiniu", New七牛KODO储存器(七牛KODO储存器配置{
		Bucket:          "testgoefun",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("qiniu_access_key"),
		AccessKeySecret: ecore.E读环境变量("qiniu_access_secret"),
		StorageConfig: &storage.Config{
			// 空间对应的机房
			Zone: &storage.Zone_as0,
			// 是否使用https域名
			UseHTTPS: false,
			// 上传是否使用CDN上传加速
			UseCdnDomains: false,
		},
		domain:  "http://rv3k84oie.sabkt.gdipper.com",
		Private: true,
	}))

	err := 文件系统.Put("a.txt", "abc")
	assert.Equal(t, nil, err, "Put 错误")

	数据, err := 文件系统.Get("a.txt")
	assert.Equal(t, nil, err, "Get 错误")
	assert.Equal(t, "abc", string(数据), "Get 错误")

	mimetype, err := 文件系统.MimeType("a.txt")
	assert.Equal(t, nil, err, "MimeType 错误")
	assert.Equal(t, "text/plain", mimetype, "MimeType 错误")

	move, err := 文件系统.Move("a.txt", "b.txt")
	assert.Equal(t, nil, err, "Move 错误")
	assert.Equal(t, true, move, "Move 错误")
	//
	copyx, err := 文件系统.Copy("b.txt", "c.txt")
	assert.Equal(t, nil, err, "Copy 错误")
	assert.Equal(t, true, copyx, "Copy 错误")
	//
	Exists, err := 文件系统.Exists("c.txt")
	assert.Equal(t, nil, err, "Exists 错误")
	assert.Equal(t, true, Exists, "Exists 错误")
	//
	size, _ := 文件系统.Size("c.txt")
	assert.Equal(t, int64(3), size, "Size 错误")
	//
	list, err := 文件系统.List("")
	assert.Equal(t, nil, err, "List 错误")
	assert.Equal(t, []string{"test/b.txt", "test/c.txt"}, list, "List 错误")
	//
	err = 文件系统.Delete("a.txt")
	// 断言 err 会出现异常
	assert.NotNil(t, err, "Delete 错误")

	err = 文件系统.Delete("b.txt")
	assert.Equal(t, nil, err, "Delete 错误")
	err = 文件系统.Delete("c.txt")
	assert.Equal(t, nil, err, "Delete 错误")

	list, err = 文件系统.List("")
	assert.Equal(t, nil, err, "List 错误")
	assert.Equal(t, []string{}, list, "List 错误")
}

func Test文件储存类(t *testing.T) {
	ecore.E加载环境变量_从文件("./.env")

	文件系统 := New文件储存类("local")
	文件系统.设置储存器("local", New本地文件储存器("./file/"))
	文件系统.设置储存器("oss", New阿里云OSS储存器(阿里云OSS储存器配置{
		Endpoint:        "oss-cn-guangzhou.aliyuncs.com",
		Bucket:          "testupload123",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("aliyun_access_key"),
		AccessKeySecret: ecore.E读环境变量("aliyun_access_secret"),
	}))
	文件系统.设置储存器("qiniu", New七牛KODO储存器(七牛KODO储存器配置{
		Bucket:          "testgoefun",
		PathPrefix:      "test/",
		AccessKeyID:     ecore.E读环境变量("qiniu_access_key"),
		AccessKeySecret: ecore.E读环境变量("qiniu_access_secret"),
		StorageConfig: &storage.Config{
			// 空间对应的机房
			Zone: &storage.Zone_as0,
			// 是否使用https域名
			UseHTTPS: false,
			// 上传是否使用CDN上传加速
			UseCdnDomains: false,
		},
		domain:  "http://rv3k84oie.sabkt.gdipper.com",
		Private: true,
	}))

	测试储存器 := []string{"local", "oss", "qiniu"}
	//测试储存器 := []string{"oss"}
	// 循环测试储存器
	for _, 储存器 := range 测试储存器 {
		文件系统.当前储存器 = 储存器
		print("当前储存器:", 储存器, "\n")

		err := 文件系统.Put("a.txt", "abc")
		assert.Equal(t, nil, err, "Put 错误")

		数据, err := 文件系统.Get("a.txt")
		assert.Equal(t, nil, err, "Get 错误")
		assert.Equal(t, "abc", string(数据), "Get 错误")

		//mimetype, err := 文件系统.MimeType("a.txt")
		//assert.Equal(t, nil, err, "MimeType 错误")
		//assert.Equal(t, "text/plain", mimetype, "MimeType 错误")

		move, err := 文件系统.Move("a.txt", "b.txt")
		assert.Equal(t, nil, err, "Move 错误")
		assert.Equal(t, true, move, "Move 错误")
		//
		copyx, err := 文件系统.Copy("b.txt", "c.txt")
		assert.Equal(t, nil, err, "Copy 错误")
		assert.Equal(t, true, copyx, "Copy 错误")
		//
		Exists, err := 文件系统.Exists("c.txt")
		assert.Equal(t, nil, err, "Exists 错误")
		assert.Equal(t, true, Exists, "Exists 错误")
		//
		size, _ := 文件系统.Size("c.txt")
		assert.Equal(t, int64(3), size, "Size 错误")
		//
		list, err := 文件系统.List("")
		assert.Equal(t, nil, err, "List 错误")
		assert.Equal(t, []string{"b.txt", "c.txt"}, list, "List 错误")
		//
		err = 文件系统.Delete("a.txt")
		// 断言 err 会出现异常
		//assert.NotNil(t, err, "Delete 错误")

		err = 文件系统.Delete("b.txt")
		assert.Equal(t, nil, err, "Delete 错误")
		err = 文件系统.Delete("c.txt")
		assert.Equal(t, nil, err, "Delete 错误")

		list, err = 文件系统.List("")
		assert.Equal(t, nil, err, "List 错误")
		assert.Equal(t, []string{}, list, "List 错误")
	}

}
