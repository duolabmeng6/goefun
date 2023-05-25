package edb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	jsonData, err := queryAndReturnJSON(db, "select id as idSuper, ctime as created_at, title as titleSuper from t1 limit 0,5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}
func TestDB2(t *testing.T) {
	op := NewMysql数据库操作类()
	err := op.E连接数据库("root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}
	defer op.E关闭数据库()

	// 执行查询操作
	queryStr := "SELECT id, name, age FROM users WHERE age >= :age"
	results, err := op.E查询(queryStr, H{"age": 20})
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}
	jsonData, _ := json.Marshal(results)
	fmt.Println(string(jsonData))

	queryStr2 := "SELECT count(id) FROM users WHERE age >= :age"
	count, err := op.Count(queryStr2, H{"age": 20})

	fmt.Println("count", count, err)

	// 执行插入操作
	//insertStr := "INSERT INTO users(name, age) VALUES(:name, :age)"
	//args := map[string]interface{}{"name": "Alice", "age": 28}
	//
	//id, err := op.Insert(insertStr, args)
	//if err != nil {
	//	fmt.Println("插入失败：", err)
	//	return
	//}
	//fmt.Println("插入成功，id为：", id)
	//
	//if err != nil {
	//	fmt.Println("开启事务失败：", err)
	//	return
	//}
	//// 执行更新操作
	//updateStr := "UPDATE users SET age = age + :v WHERE name = :name"
	//args = map[string]interface{}{
	//	"name": "Alice",
	//	"v":    2,
	//}
	//
	//affected, err := op.Update(updateStr, args)
	//if err != nil {
	//	fmt.Println("更新失败：", err)
	//	return
	//}
	//fmt.Println("更新成功，影响行数为：", affected)

	// 执行删除操作
	//deleteStr := "DELETE FROM users WHERE age > :age"
	//affected, err := op.Delete(deleteStr, map[string]interface{}{
	//	"age": 30,
	//})
	//if err != nil {
	//	fmt.Println("删除失败：", err)
	//}
	//fmt.Println("删除成功，影响行数为：", affected)

}

func TestDBAutoPage(t *testing.T) {
	fmt.Println("TestDBAutoPage")
	op := NewMysql数据库操作类()
	err := op.E连接数据库("root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}
	defer op.E关闭数据库()

	page := 1
	pageSize := 10

	queryStr := "SELECT id, name, age FROM users WHERE age >= :age LIMIT :limit,:pageSize"
	results, err := op.E查询(queryStr, H{
		"age":      20,
		"limit":    page,
		"pageSize": pageSize,
	})
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}
	jsonData, _ := json.Marshal(results)
	fmt.Println(string(jsonData))

	queryStr2 := "SELECT COUNT(*) FROM users WHERE age >= :age"
	count, err := op.Count(queryStr2, H{"age": 20})

	fmt.Println("count", count, err)

}
func TestDB查询构建器(t *testing.T) {
	type ArticlesIndexRequest struct {
		Keywords string `i:"keywords"`
		PerPage  int64  `i:"perPage" rule:"required" msg:"PerPage 必填"`
		Page     int64  `i:"page" rule:"required" msg:"Page 必填"`
		OrderBy  string `i:"orderBy" default:"id"`
		OrderDir string `i:"orderDir" default:"desc"`
	}
	//SELECT id,title,content,created_at,updated_at FROM articles ORDER BY id desc, title desc LIMIT 0,10
	var req ArticlesIndexRequest
	req.Keywords = "a"
	req.PerPage = 10
	req.Page = 1
	req.OrderBy = "id"
	req.OrderDir = "desc"
	//Where(map[string]interface{}{
	//	"title like":   "%" + req.Keywords + "%",
	//	"id >":         1,
	//	"created_at <": "2020-01-01 00:00:00",
	//	"content like": "%apple%",
	//}).
	dialect := NewMySQLQueryBuilder()
	qb := dialect.From("articles").
		Select("id", "title", "content", "created_at", "updated_at").
		Where("title", "like", "%"+req.Keywords+"%").
		//Where("id", "=", 2146).
		OrderBy(req.OrderBy, req.OrderDir).
		Limit(int(req.Page), int(req.PerPage))

	query, param := qb.ToSQL()
	query2, param := qb.Count().ToSQL()

	ecore.E调试输出(query, param)
	ecore.E调试输出(query2)

	op := NewMysql数据库操作类()
	op.E连接数据库("root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local")
	//result, err := op.QueryRaw(query, param)
	//ecore.E调试输出(err, result)

	count, _ := op.CountRaw(query2, param)
	ecore.E调试输出(count)

	//s, _ := json.MarshalIndent(result, "", "  ")
	//fmt.Println(string(s))

}
