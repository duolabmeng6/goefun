package ejson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type EJsonI interface {
	E加载(字符串 string) error
	E导出为JSON() (string, error)
	E置值(pathKey string, value any) error
	E取值(pathKey string) (any, error)
}

type EJson struct {
	data map[string]interface{}
}

func NewEJson() *EJson {
	return &EJson{
		data: make(map[string]interface{}),
	}
}

func (e *EJson) E加载(字符串 string) error {
	return json.Unmarshal([]byte(字符串), &e.data)
}

func (e *EJson) E导出为JSON() (string, error) {
	bytes, err := json.Marshal(e.data)
	return string(bytes), err
}

func (e *EJson) E置值(pathKey string, value any) error {
	if pathKey == "" {
		return errors.New("pathKey cannot be empty")
	}

	keys := strings.Split(pathKey, ".")
	lastIndex := len(keys) - 1
	currentMap := e.data

	for i, key := range keys {
		if i == lastIndex {
			currentMap[key] = value
			return nil
		}

		if nextMap, ok := currentMap[key].(map[string]interface{}); ok {
			currentMap = nextMap
		} else {
			// Create a new map if the key does not exist
			newMap := make(map[string]interface{})
			currentMap[key] = newMap
			currentMap = newMap
		}
	}

	return nil
}

func (e *EJson) E取值(pathKey string) (any, error) {
	// 将 [0] 转换为 .0
	re := regexp.MustCompile(`\[(\d+)\]`)
	pathKey = re.ReplaceAllString(pathKey, ".$1")

	keys := strings.Split(pathKey, ".")
	var current interface{} = e.data

	for _, key := range keys {
		currentVal := reflect.ValueOf(current)

		switch currentVal.Kind() {
		case reflect.Slice, reflect.Array:
			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, fmt.Errorf("invalid array index '%s'", key)
			}
			if idx < 0 || idx >= currentVal.Len() {
				return nil, fmt.Errorf("array index '%d' out of range", idx)
			}
			current = currentVal.Index(idx).Interface()

		case reflect.Map:
			currentMap, ok := current.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("not a map at '%s'", key)
			}
			current, ok = currentMap[key]
			if !ok {
				return nil, fmt.Errorf("key '%s' not found", key)
			}

		default:
			return nil, fmt.Errorf("invalid type at '%s', expected map or array", key)
		}
	}

	return current, nil
}
