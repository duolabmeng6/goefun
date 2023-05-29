package edb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MySQLQueryBuilder struct {
	db            *Mysql数据库操作类
	selectCols    []string
	fromTable     string
	whereConds    []string
	whereArgs     []interface{}
	orWhereConds  []string
	orWhereArgs   []interface{}
	orderCols     []string
	limit         []int64
	isDeleteQuery bool
}

func NewMySQLQueryBuilder() *MySQLQueryBuilder {
	return &MySQLQueryBuilder{}
}

func NewMySQL查询构建器(db *Mysql数据库操作类) *MySQLQueryBuilder {
	return &MySQLQueryBuilder{
		db: db,
	}
}

func (qb *MySQLQueryBuilder) Select(cols ...string) *MySQLQueryBuilder {
	qb.selectCols = cols
	return qb
}

func (qb *MySQLQueryBuilder) From(table string) *MySQLQueryBuilder {
	qb.fromTable = table
	//执行后清空所有的条件
	qb.whereConds = nil
	qb.whereArgs = nil
	qb.orWhereConds = nil
	qb.orWhereArgs = nil
	qb.orderCols = nil
	qb.limit = nil

	return qb
}

func (qb *MySQLQueryBuilder) Where(column string, operator string, value interface{}) *MySQLQueryBuilder {
	qb.whereConds = append(qb.whereConds, fmt.Sprintf("%s %s ?", column, operator))
	qb.whereArgs = append(qb.whereArgs, value)
	return qb
}
func (qb *MySQLQueryBuilder) OrWhere(column string, operator string, value interface{}) *MySQLQueryBuilder {
	qb.orWhereConds = append(qb.orWhereConds, fmt.Sprintf("%s %s ?", column, operator))
	qb.orWhereArgs = append(qb.orWhereArgs, value)
	return qb
}
func (qb *MySQLQueryBuilder) OrderBy(cols ...string) *MySQLQueryBuilder {
	qb.orderCols = cols
	return qb
}

func (qb *MySQLQueryBuilder) Limit(offset, limit int64) *MySQLQueryBuilder {
	qb.limit = []int64{offset, limit}
	return qb
}

func (qb *MySQLQueryBuilder) Paginate(total int64, current_pages int64) *MySQLQueryBuilder {
	if current_pages <= 0 {
		current_pages = 1
	}
	current_pages = (current_pages - 1) * total

	qb.limit = []int64{current_pages, total}
	return qb
}
func (qb *MySQLQueryBuilder) Count() (int64, error) {
	//copy一个新的对象返回
	countQueryBuilder := NewMySQL查询构建器(qb.db)
	countQueryBuilder.selectCols = []string{"count(*) as count"}
	countQueryBuilder.fromTable = qb.fromTable
	countQueryBuilder.whereArgs = qb.whereArgs
	countQueryBuilder.whereConds = qb.whereConds
	countQueryBuilder.orWhereArgs = qb.orWhereArgs
	countQueryBuilder.orWhereConds = qb.orWhereConds
	return countQueryBuilder.GetInt()
}

func (qb *MySQLQueryBuilder) ToSQL() (string, []interface{}) {
	var sb strings.Builder
	args := make([]interface{}, 0, len(qb.selectCols)+len(qb.whereArgs)+len(qb.orWhereArgs)+len(qb.orderCols)+2)

	if qb.isDeleteQuery {
		sb.WriteString("DELETE FROM ")
		sb.WriteString(qb.fromTable)
	} else {
		sb.WriteString("SELECT ")
		if len(qb.selectCols) == 0 {
			sb.WriteString("*")
		} else {
			sb.WriteString(strings.Join(qb.selectCols, ", "))
		}

		sb.WriteString(" FROM ")
		sb.WriteString(qb.fromTable)
	}

	if len(qb.whereConds) > 0 || len(qb.orWhereConds) > 0 {
		sb.WriteString(" WHERE ")

		// Join WHERE conditions
		if len(qb.whereConds) > 0 {
			sb.WriteString(strings.Join(qb.whereConds, " AND "))
			args = append(args, qb.whereArgs...)
		}

		// Join OR WHERE conditions
		if len(qb.orWhereConds) > 0 {
			if len(qb.whereConds) > 0 {
				sb.WriteString(" OR ")
			}
			sb.WriteString(strings.Join(qb.orWhereConds, " OR "))
			args = append(args, qb.orWhereArgs...)
		}
	}

	if !qb.isDeleteQuery && len(qb.orderCols) > 0 {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(strings.Join(qb.orderCols, " "))
	}

	if !qb.isDeleteQuery && len(qb.limit) > 0 {
		sb.WriteString(" LIMIT " + strconv.FormatInt(qb.limit[0], 10) + "," + strconv.FormatInt(qb.limit[1], 10))
	}

	query := sb.String()
	return query, args
}

func (qb *MySQLQueryBuilder) Get() ([]map[string]interface{}, error) {
	query, args := qb.ToSQL()
	fmt.Println("query", query)
	fmt.Println("args", args)
	rows, err := qb.db.QueryRaw(query, args)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (qb *MySQLQueryBuilder) GetInt() (int64, error) {
	query, args := qb.ToSQL()
	fmt.Println("query", query)
	fmt.Println("args", args)
	rows, err := qb.db.CountRaw(query, args)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (qb *MySQLQueryBuilder) First() (map[string]interface{}, error) {
	//设置limit=1
	qb.limit = []int64{0, 1}
	query, args := qb.ToSQL()
	fmt.Println("query", query)
	fmt.Println("args", args)
	result, err := qb.db.QueryRaw(query, args)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("not found")
	}
	return result[0], nil
}

func (qb *MySQLQueryBuilder) Delete() (int64, error) {
	qb.isDeleteQuery = true
	query, args := qb.ToSQL()
	qb.isDeleteQuery = false
	fmt.Println("query", query, "args", args)
	i, err := qb.db.ExecRaw(query, args)

	if err != nil {
		return 0, err
	}
	return i, nil
}

func (qb *MySQLQueryBuilder) Insert(data map[string]interface{}) (int64, error) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	for column, value := range data {
		columns = append(columns, column)
		values = append(values, value)
	}

	columnStr := strings.Join(columns, ", ")
	valueStr := strings.Repeat("?, ", len(columns))
	valueStr = valueStr[:len(valueStr)-2]

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", qb.fromTable, columnStr, valueStr)
	fmt.Println("INSERT", query)
	stmt, err := qb.db.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (qb *MySQLQueryBuilder) Update(data map[string]interface{}) (int64, error) {
	if len(data) == 0 {
		return 0, errors.New("no data provided for update")
	}

	setValues := make([]string, 0)
	setArgs := make([]interface{}, 0)

	for column, value := range data {
		setValues = append(setValues, fmt.Sprintf("%s = ?", column))
		setArgs = append(setArgs, value)
	}

	setStr := strings.Join(setValues, ", ")

	whereStr, whereArgs := qb.getWhereClause()

	args := append(setArgs, whereArgs...)

	query := fmt.Sprintf("UPDATE %s SET %s%s", qb.fromTable, setStr, whereStr)

	fmt.Println("UPDATE", query)
	fmt.Println("args", args)

	stmt, err := qb.db.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (qb *MySQLQueryBuilder) getWhereClause() (string, []interface{}) {
	whereClause := ""
	args := make([]interface{}, 0)

	if len(qb.whereConds) > 0 || len(qb.orWhereConds) > 0 {
		whereClause += " WHERE "

		// Join WHERE conditions
		if len(qb.whereConds) > 0 {
			whereClause += strings.Join(qb.whereConds, " AND ")
			args = append(args, qb.whereArgs...)
		}

		// Join OR WHERE conditions
		if len(qb.orWhereConds) > 0 {
			if len(qb.whereConds) > 0 {
				whereClause += " OR "
			}
			whereClause += strings.Join(qb.orWhereConds, " OR ")
			args = append(args, qb.orWhereArgs...)
		}
	}

	return whereClause, args
}
