package ecache

// 缓存操作接口
type 缓存接口 interface {
	Set(key string, value interface{}, 倒计时秒数 int64) error
	Get(key string) (interface{}, error)
	Del(key string) error
}

type 缓存数据 struct {
	DataType interface{} `json:"DataType"`
	Data     interface{} `json:"Data"`
	EndTime  int64       `json:"EndTime"`
}
