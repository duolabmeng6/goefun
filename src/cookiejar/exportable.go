package cookiejar

import (
	"encoding/json"
)

type Entry = entry

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

func (j *Jar) Entries() map[string]map[string]Entry {
	return j.entries
}

func (j *Jar) SetEntries(entries map[string]map[string]Entry) {
	j.entries = entries
}

func (j *Jar) JsonSerialize() ([]byte, error) {
	return json.Marshal(j.Entries())
}

func (j *Jar) JsonDeserialize(data []byte) error {
	entries := make(map[string]map[string]Entry, 0)
	err := json.Unmarshal(data, &entries)
	if err != nil {
		return err
	}
	j.SetEntries(entries)
	return nil
}
