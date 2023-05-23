package efile

// 文件操作接口
type 文件储存接口 interface {
	Put(key string, data interface{}) error      // 保存文件
	Get(key string) ([]byte, error)              // 获取文件
	Delete(key string) error                     // 删除文件
	Move(key string, tokey string) (bool, error) // 移动文件
	Copy(key string, tokey string) (bool, error) // 复制文件
	Exists(key string) (bool, error)             // 文件是否存在
	Size(key string) (int64, error)              // 文件大小
	List(key string) ([]string, error)           // 文件列表
	MimeType(key string) (string, error)         // 文件Mime
}
