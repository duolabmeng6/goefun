// Package cookiejar Cookie 管理扩展方法（序列化/反序列化等）
package cookiejar

import (
	"encoding/json"
)

// Entry 是内部 entry 类型的导出别名，便于外部直接访问 Cookie 条目。
type Entry = entry

// Get 根据 Cookie 名称返回对应的值，若不存在则返回空字符串。
func (j *Jar) Get(key string) string {
	for _, v := range j.entries {
		for _, vv := range v {
			if key == vv.Name {
				return vv.Value
			}
		}
	}
	return ""
}

// Entries 返回内部存储的所有 Cookie 条目的映射。
func (j *Jar) Entries() map[string]map[string]Entry {
	return j.entries
}

// SetEntries 使用给定的条目映射替换内部存储。
func (j *Jar) SetEntries(entries map[string]map[string]Entry) {
	j.entries = entries
}

// JsonSerialize 将所有 Cookie 条目序列化为 JSON 字节切片。
func (j *Jar) JsonSerialize() ([]byte, error) {
	return json.Marshal(j.Entries())
}

// JsonDeserialize 从 JSON 字节切片反序列化并设置 Cookie 条目。
func (j *Jar) JsonDeserialize(data []byte) error {
	entries := make(map[string]map[string]Entry, 0)
	err := json.Unmarshal(data, &entries)
	if err != nil {
		return err
	}
	j.SetEntries(entries)
	return nil
}
