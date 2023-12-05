// Package esqlite Description: sqlite数据库操作封装
package esqlite

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type ESqliteI interface {
	E打开数据库(数据库文件 string) error
	E打开内存数据库() error
	E关闭数据库() error
	E执行SQL(命令 string) error
	E执行查询SQL(命令 string) ([]map[string]interface{}, error)
	E执行事务(f func(tx *gorm.DB) error) error
}

type ESqlite struct {
	ESqliteI
	db *gorm.DB
}

func NewESqlite() *ESqlite {
	return &ESqlite{
		db: nil,
	}
}

func (e *ESqlite) E打开数据库(数据库文件 string) error {
	var err error
	e.db, err = gorm.Open(sqlite.Open(数据库文件), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("无法打开数据库: %w", err)
	}
	return nil
}

func (e *ESqlite) E打开内存数据库() error {
	var err error
	e.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("无法打开内存数据库: %w", err)
	}
	return nil
}

func (e *ESqlite) E关闭数据库() error {
	var err error
	sqlDB, err := e.db.DB()
	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("无法关闭数据库: %w", err)
	}
	return nil
}

func (e *ESqlite) E执行SQL(命令 string) error {
	result := e.db.Exec(命令)
	if result.Error != nil {
		return fmt.Errorf("执行SQL失败: %w", result.Error)
	}
	return nil
}
func (e *ESqlite) E执行查询SQL(命令 string) ([]map[string]interface{}, error) {
	rows, err := e.db.Raw(命令).Rows()
	if err != nil {
		return nil, fmt.Errorf("执行查询SQL失败: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("获取列名失败: %w", err)
	}

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	results := []map[string]interface{}{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, fmt.Errorf("扫描行失败: %w", err)
		}

		resultMap := map[string]interface{}{}
		for k, v := range values {
			key := columns[k]
			value := ""
			switch v := v.(type) {
			case []byte:
				value = string(v)
			default:
				value = fmt.Sprintf("%v", v)
			}
			resultMap[key] = value
		}
		results = append(results, resultMap)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("读取行失败: %w", err)
	}

	return results, nil
}

func (e *ESqlite) E执行事务(f func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) error {
		return f(tx)
	})
}
