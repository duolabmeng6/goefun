package edb

import (
	"fmt"
	"strconv"
	"strings"
)

type MySQLQueryBuilder struct {
	selectCols []string
	fromTable  string
	whereConds []string
	whereArgs  []interface{}
	orderCols  []string
	limit      []int
}

func NewMySQLQueryBuilder() *MySQLQueryBuilder {
	return &MySQLQueryBuilder{}
}

func NewMySQL查询构建器() *MySQLQueryBuilder {
	return &MySQLQueryBuilder{}
}

func (qb *MySQLQueryBuilder) Select(cols ...string) *MySQLQueryBuilder {
	qb.selectCols = cols
	return qb
}

func (qb *MySQLQueryBuilder) From(table string) *MySQLQueryBuilder {
	qb.fromTable = table
	return qb
}

func (qb *MySQLQueryBuilder) Where(column string, operator string, value interface{}) *MySQLQueryBuilder {
	qb.whereConds = append(qb.whereConds, fmt.Sprintf("%s %s ?", column, operator))
	qb.whereArgs = append(qb.whereArgs, value)
	return qb
}

func (qb *MySQLQueryBuilder) OrderBy(cols ...string) *MySQLQueryBuilder {
	qb.orderCols = cols
	return qb
}

func (qb *MySQLQueryBuilder) Limit(offset, limit int) *MySQLQueryBuilder {
	//如果 offset = 1 需要 offset-1
	if offset > 0 {
		offset = offset - 1
	}
	qb.limit = []int{offset, limit}
	return qb
}
func (qb *MySQLQueryBuilder) Count() *MySQLQueryBuilder {
	countQueryBuilder := &MySQLQueryBuilder{
		selectCols: []string{"COUNT(*)"},
		fromTable:  qb.fromTable,
		whereConds: make([]string, len(qb.whereConds)),
		whereArgs:  make([]interface{}, len(qb.whereArgs)),
	}

	copy(countQueryBuilder.whereConds, qb.whereConds)
	copy(countQueryBuilder.whereArgs, qb.whereArgs)

	return countQueryBuilder
}

func (qb *MySQLQueryBuilder) ToSQL() (string, []interface{}) {
	var sb strings.Builder
	args := make([]interface{}, 0, len(qb.selectCols)+len(qb.whereArgs)+len(qb.orderCols)+2)

	sb.WriteString("SELECT ")
	if len(qb.selectCols) == 0 {
		sb.WriteString("*")
	} else {
		sb.WriteString(strings.Join(qb.selectCols, ", "))
	}

	sb.WriteString(" FROM ")
	sb.WriteString(qb.fromTable)

	if len(qb.whereConds) > 0 {
		sb.WriteString(" WHERE ")
		sb.WriteString(strings.Join(qb.whereConds, " AND "))
		args = append(args, qb.whereArgs...)
	}

	if len(qb.orderCols) > 0 {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(strings.Join(qb.orderCols, " "))
	}

	if len(qb.limit) > 0 {
		sb.WriteString(" LIMIT " + strconv.Itoa(qb.limit[0]) + "," + strconv.Itoa(qb.limit[1]))
	}

	query := sb.String()
	return query, args
}
