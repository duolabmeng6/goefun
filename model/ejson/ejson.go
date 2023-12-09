// Package ejson Description: 用于处理json的工具
package ejson

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type EJsonI interface {
	E加载(字符串 string) error
	E加载从env配置内容(字符串 string) error
	E加载从ini配置内容(字符串 string) error
	E导出为JSON() (string, error)
	E置值(pathKey string, value any) error
	E取值(pathKey string) (any, error)
}

type EJson struct {
	EJsonI
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

// E加载从env配置内容 从.env格式字符串加载内容
func (e *EJson) E加载从env配置内容(env内容 string) error {
	if env内容 == "" {
		return errors.New("env内容不能为空")
	}

	scanner := bufio.NewScanner(strings.NewReader(env内容))
	result := make(map[string]interface{})

	for scanner.Scan() {
		line := scanner.Text()

		// 跳过空行和注释
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return errors.New("env格式错误: " + line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 简单处理值的引号
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		} else if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			value = strings.Trim(value, "'")
		}

		result[key] = value
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	e.data = result
	return nil
}

// E加载从ini配置内容 从INI格式字符串加载内容
func (e *EJson) E加载从ini配置内容(INI字符串 string) error {
	cfg, err := ini.Load([]byte(INI字符串))
	if err != nil {
		return fmt.Errorf("读取INI字符串失败: %w", err)
	}

	result := make(map[string]interface{})

	for _, section := range cfg.Sections() {
		sectionName := section.Name()
		if sectionName == ini.DefaultSection {
			continue
		}

		sectionMap := make(map[string]interface{})
		for _, key := range section.Keys() {
			sectionMap[key.Name()] = key.Value()
		}

		result[sectionName] = sectionMap
	}

	e.data = result
	return nil
}
