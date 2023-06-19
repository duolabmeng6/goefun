// Package edb 数据库操作
package edb

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/elliotchance/orderedmap/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"time"
)

type DatabaseOperator interface {
	Connect(string) error                                 // 返回连接数据库时可能出现的错误 中文函数名: 连接数据库
	Query(string) ([]map[string]interface{}, error)       // 接收查询语句作为参数，返回查询结果的列表和可能出现的错误 中文函数名: 查询
	Insert(string, map[string]interface{}) (int64, error) // 接收插入语句和待插入数据作为参数，返回插入操作对应的id值和可能出现的错误 中文函数名: 插入
	Update(string, map[string]interface{}) (int64, error) // 接收更新语句作为参数，返回更新操作对应的行数和可能出现的错误 中文函数名: 更新
	Delete(string, map[string]interface{}) (int64, error) // 接收删除语句作为参数，返回删除操作对应的行数和可能出现的错误 中文函数名: 删除
	Count(string, map[string]interface{}) (int64, error)  // 接收查询语句作为参数，返回查询结果的数量和可能出现的错误 中文函数名: 总数
	BeginTransaction() error                              // 返回开始事务时可能出现的错误 中文函数名: 开始事务
	CommitTransaction() error                             // 返回提交事务时可能出现的错误 中文函数名: 提交事务
	RollbackTransaction() error                           // 返回回滚事务时可能出现的错误 中文函数名: 回滚事务
	QueryWithPagination(string, int, int) ([]map[string]interface{}, error)
}

// H 简化map[string]interface{}的写法 edb.H{"name": "张三", "age": 18} 用起来方便一点
type H map[string]any

type Mysql数据库操作类 struct {
	db     *sqlx.DB
	tx     *sqlx.Tx
	dbName string
}

func NewMysql数据库操作类() *Mysql数据库操作类 {
	return &Mysql数据库操作类{}
}

// E关闭连接
func (op *Mysql数据库操作类) E关闭数据库() {
	if op.db != nil {
		err := op.db.Close()
		if err != nil {
			panic("关闭数据库连接失败：" + err.Error())
			return
		}
	}
}

// E连接数据库 用于连接数据库，返回连接数据库时可能出现的错误
// 参数datasourceName是数据库连接信息，格式为"用户名:密码@tcp(IP:端口)/数据库名?charset=utf8"
func (op *Mysql数据库操作类) E连接数据库(数据库连接信息 string) error {
	return op.Connect(数据库连接信息)
}

// E查询 用于执行查询操作，返回查询结果的列表和可能出现的错误
// 参数queryStr是查询语句，参数args是查询语句中的参数，格式为map[string]interface{}{"参数名": 参数值}
func (op *Mysql数据库操作类) E查询(查询语句 string, 参数 map[string]interface{}) ([]map[string]interface{}, error) {
	return op.Query(查询语句, 参数)
}

// E插入 用于执行插入操作，返回插入操作对应的id值和可能出现的错误
// 参数insertStr是插入语句，参数args是插入语句中的参数，格式为map[string]interface{}{"参数名": 参数值}
func (op *Mysql数据库操作类) E插入(插入语句 string, 参数 map[string]interface{}) (int64, error) {
	return op.Insert(插入语句, 参数)
}

// E更新 用于执行更新操作，返回更新操作对应的行数和可能出现的错误
// 参数updateStr是更新语句，参数args是更新语句中的参数，格式为map[string]interface{}{"参数名": 参数值}
func (op *Mysql数据库操作类) E更新(更新语句 string, 参数 map[string]interface{}) (int64, error) {
	return op.Update(更新语句, 参数)
}

// E删除 用于执行删除操作，返回删除操作对应的行数和可能出现的错误
// 参数deleteStr是删除语句，参数args是删除语句中的参数，格式为map[string]interface{}{"参数名": 参数值}
func (op *Mysql数据库操作类) E删除(删除语句 string, 参数 map[string]interface{}) (int64, error) {
	return op.Delete(删除语句, 参数)
}

// E开始事务 用于开始事务，返回开始事务时可能出现的错误
func (op *Mysql数据库操作类) E开始事务() error {
	return op.BeginTransaction()
}

// E提交事务 用于提交事务，返回提交事务时可能出现的错误
func (op *Mysql数据库操作类) E提交事务() error {
	return op.CommitTransaction()
}

// E回滚事务 用于回滚事务，返回回滚事务时可能出现的错误
func (op *Mysql数据库操作类) E回滚事务() error {
	return op.RollbackTransaction()
}

func (op *Mysql数据库操作类) Connect(datasourceName string) error {
	var err error
	op.db, err = sqlx.Connect("mysql", datasourceName)
	if err != nil {
		return err
	}
	// "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local"
	//分析datasourceName的数据获取数据库的名称 gotest
	//获取第一个/的位置
	op.dbName = ecore.StrCut(datasourceName, "/$?")

	return nil
}

// LIMIT
// 用于拼接LIMIT语句，返回拼接后的LIMIT语句
// Page 第几页
// PerPage 每页多少条
func LIMIT(Page interface{}, PerPage interface{}) string {
	//转换为 int64
	page, _ := strconv.ParseInt(fmt.Sprintf("%v", Page), 10, 64)
	perPage, _ := strconv.ParseInt(fmt.Sprintf("%v", PerPage), 10, 64)

	return fmt.Sprintf(" LIMIT %d,%d", (page-1)*perPage, PerPage)
}

// OrderBY 接收任意参数组合 用于拼接ORDER BY语句，返回拼接后的ORDER BY语句
// 调用方式 OrderBY("id", "desc", "title", "desc")
func OrderBY(参数 ...interface{}) string {
	var orderStr string
	for i := 0; i < len(参数); i++ {
		if i%2 == 0 {
			orderStr += fmt.Sprintf(" %v", 参数[i])
		} else {
			orderStr += fmt.Sprintf(" %v,", 参数[i])
		}
	}
	return " ORDER BY" + strings.TrimRight(orderStr, ",")

}

// QueryRaw 是使用 ? 占位符的
func (op *Mysql数据库操作类) QueryRaw(queryStr string, args []interface{}) ([]map[string]interface{}, error) {
	stmt, err := op.db.Prepare(queryStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	for i, columnType := range columnTypes {
		switch columnType.DatabaseTypeName() {
		case "INT", "BIGINT", "UNSIGNED INT", "UNSIGNED BIGINT":
			values[i] = new(int)
		case "FLOAT", "DOUBLE", "DECIMAL", "NUMERIC":
			values[i] = new(float64)
		case "DATE", "DATETIME", "TIMESTAMP":
			values[i] = new(sql.NullTime)
		case "TEXT", "VARCHAR":
			values[i] = new(string)
		case "BIT":
			values[i] = new(bool)
		case "BLOB":
			values[i] = &[]byte{}
		default:
			//打印未能处理的数据类型
			fmt.Println(errors.New("Query 未能正确处理的数据类型: " + columnType.DatabaseTypeName()))
			values[i] = new(string)
		}
	}
	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, value := range values {
			switch v := value.(type) {
			case int, float64, bool, []byte:
				result[columns[i]] = v
			case string:
				result[columns[i]] = v
			case time.Time:
				if columnTypes[i].DatabaseTypeName() == "DATE" {
					result[columns[i]] = v.Format("2006-01-02")
				} else {
					result[columns[i]] = v.Format(time.RFC3339Nano)
				}
			case sql.NullTime:
				if v.Valid {
					result[columns[i]] = v.Time.Format(time.RFC3339Nano)
				} else {
					result[columns[i]] = nil
				}
			case *int:
				if v != nil {
					result[columns[i]] = *v
				} else {
					result[columns[i]] = nil
				}
			case *float64:
				if v != nil {
					result[columns[i]] = *v
				} else {
					result[columns[i]] = nil
				}
			case *bool:
				if v != nil {
					result[columns[i]] = *v
				} else {
					result[columns[i]] = nil
				}
			case *[]byte:
				if v != nil {
					result[columns[i]] = *v
				} else {
					result[columns[i]] = nil
				}
			case *string:
				if v != nil {
					result[columns[i]] = *v
				} else {
					result[columns[i]] = nil
				}
			case *time.Time:
				if v != nil {
					if columnTypes[i].DatabaseTypeName() == "DATE" {
						result[columns[i]] = v.Format("2006-01-02")
					} else {
						result[columns[i]] = v.Format(time.RFC3339Nano)
					}
				} else {
					result[columns[i]] = nil
				}
			case *sql.NullTime:
				if v != nil {
					if v.Valid {
						result[columns[i]] = v.Time.Format(time.RFC3339Nano)
					} else {
						result[columns[i]] = nil
					}
				} else {
					result[columns[i]] = nil
				}
			}
		}

		results = append(results, result)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return results, nil
}

func (op *Mysql数据库操作类) Query(queryStr string, args map[string]interface{}) ([]map[string]interface{}, error) {
	rows, err := op.db.NamedQuery(queryStr, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	for i, columnType := range columnTypes {
		switch columnType.DatabaseTypeName() {
		case "INT", "BIGINT", "UNSIGNED INT", "UNSIGNED BIGINT":
			values[i] = new(int)
		case "FLOAT", "DOUBLE", "DECIMAL", "NUMERIC":
			values[i] = new(float64)
		case "DATE", "DATETIME", "TIMESTAMP":
			values[i] = new(time.Time)
		case "TEXT", "VARCHAR":
			values[i] = new(string)
		case "BIT":
			values[i] = new(bool)
		case "BLOB":
			values[i] = &[]byte{}
		default:
			//打印未能处理的数据类型
			fmt.Println(errors.New("Query 未能正确处理的数据类型: " + columnType.DatabaseTypeName()))
			values[i] = new(string)
		}
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, value := range values {
			switch v := value.(type) {
			case *int, *float64, *bool, *[]byte:
				result[columns[i]] = v
			case *string:
				result[columns[i]] = *v
			case *time.Time:
				if columnTypes[i].DatabaseTypeName() == "DATE" {
					result[columns[i]] = v.Format("2006-01-02")
				} else {
					result[columns[i]] = v.Format(time.RFC3339Nano)
				}
			}
		}
		results = append(results, result)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	//
	//jsonData, err := json.Marshal(results)
	//if err != nil {
	//	return nil, err
	//}

	return results, nil
}

func (op *Mysql数据库操作类) Insert(queryStr string, args map[string]interface{}) (int64, error) {
	stmt, err := op.db.PrepareNamed(queryStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (op *Mysql数据库操作类) Update(queryStr string, args map[string]interface{}) (int64, error) {
	res, err := op.db.NamedExec(queryStr, args)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

func (op *Mysql数据库操作类) Delete(queryStr string, args map[string]interface{}) (int64, error) {
	res, err := op.db.NamedExec(queryStr, args)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

func (op *Mysql数据库操作类) ExecRaw(queryStr string, args []interface{}) (int64, error) {
	stmt, err := op.db.Prepare(queryStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

// Count
func (op *Mysql数据库操作类) Count(queryStr string, args map[string]interface{}) (int64, error) {
	// 检查查询语句中是否有 COUNT 函数
	if !strings.Contains(strings.ToUpper(queryStr), "COUNT(") {
		return 0, errors.New("你的查询语句缺少 COUNT 函数")
	}
	rows, err := op.db.NamedQuery(queryStr, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := int64(0)
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	if rows.Err() != nil {
		return 0, rows.Err()
	}

	return count, nil
}

// CountRaw
func (op *Mysql数据库操作类) CountRaw(queryStr string, args []interface{}) (int64, error) {
	// 检查查询语句中是否有 COUNT 函数
	if !strings.Contains(strings.ToUpper(queryStr), "COUNT(") {
		return 0, errors.New("你的查询语句缺少 COUNT 函数")
	}
	stmt, err := op.db.Prepare(queryStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := int64(0)
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	if rows.Err() != nil {
		return 0, rows.Err()
	}

	return count, nil
}

func (op *Mysql数据库操作类) BeginTransaction() error {
	tx, err := op.db.Beginx()
	if err != nil {
		return err
	}
	op.tx = tx
	return nil
}

func (op *Mysql数据库操作类) CommitTransaction() error {
	if op.tx == nil {
		return errors.New("Transaction not found")
	}
	err := op.tx.Commit()
	op.tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (op *Mysql数据库操作类) RollbackTransaction() error {
	if op.tx == nil {
		return errors.New("Transaction not found")
	}
	err := op.tx.Rollback()
	op.tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (op *Mysql数据库操作类) GetTableInfo(table string) (*orderedmap.OrderedMap[string, map[string]interface{}], error) {
	//获取mysql数据库的表结构和备注
	query := fmt.Sprintf("SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", op.dbName, table)
	println(query)
	rows, err := op.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := orderedmap.NewOrderedMap[string, map[string]interface{}]()
	for rows.Next() {
		var column, dataType, comment string
		err := rows.Scan(&column, &dataType, &comment)
		if err != nil {
			return nil, err
		}
		columns.Set(column, map[string]interface{}{
			"dataType": dataType,
			"comment":  comment,
		})
	}
	return columns, nil
}

func (op *Mysql数据库操作类) GetAllTableName() ([]string, error) {
	// 获取mysql数据库中的所有表名称 返回 tableName
	query := fmt.Sprintf("SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = '%s'", op.dbName)
	println(query)
	rows, err := op.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]string, 0)
	for rows.Next() {
		var TABLE_NAME string
		err := rows.Scan(&TABLE_NAME)
		if err != nil {
			return nil, err
		}
		tables = append(tables, TABLE_NAME)
	}
	return tables, nil
}

func queryAndReturnJSON(db *sql.DB, query string) ([]byte, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	for i, columnType := range columnTypes {
		switch columnType.DatabaseTypeName() {
		case "INT", "BIGINT":
			values[i] = new(int)
		case "FLOAT", "DOUBLE", "DECIMAL", "NUMERIC":
			values[i] = new(float64)
		case "DATE", "DATETIME":
			values[i] = new(time.Time)
		case "TEXT", "VARCHAR":
			values[i] = new(string)
		case "BIT":
			values[i] = new(bool)
		case "BLOB":
			values[i] = &[]byte{}
		default:
			values[i] = new(string)
		}
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, value := range values {
			switch v := value.(type) {
			case *int, *float64, *bool, *[]byte:
				result[columns[i]] = v
			case *string:
				result[columns[i]] = *v
			case *time.Time:
				if columnTypes[i].DatabaseTypeName() == "DATE" {
					result[columns[i]] = v.Format("2006-01-02")
				} else {
					result[columns[i]] = v.Format(time.RFC3339Nano)
				}
			}
		}
		results = append(results, result)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
