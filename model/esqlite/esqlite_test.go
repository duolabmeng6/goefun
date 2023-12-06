package esqlite

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/model/eval"
	"github.com/glebarez/sqlite"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"testing"
)

func TestNewESqlite(t *testing.T) {
	db := NewESqlite()
	err := db.E打开数据库(ecore.E取运行源文件路径() + "/test.db")
	if err != nil {
		t.Error(err)
		return
	}
	println("打开成功")
	//err = db.E执行SQL("create table if not exists test(id int,name text)")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//println("创建成功")
	//err = db.E执行SQL("insert into test(id,name) values(1,'张三')")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//println("插入成功")
	//err = db.E执行SQL("insert into test(id,name) values(2,'李四')")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

	//查询
	rows, err := db.E执行查询SQL("select * from test")
	if err != nil {
		t.Error(err)
		return
	}
	for _, row := range rows {
		println(eval.E到文本(row))
		println(eval.E到文本(row["name"]))
	}

}

func TestNewESqlite2(t *testing.T) {
	gdb, _ := gorm.Open(sqlite.Open(ecore.E取运行源文件路径()+"/test.db"), &gorm.Config{})
	gsqldb, _ := gdb.DB()
	db := sqlx.NewDb(gsqldb, "sqlite3")

	rows, err := db.Query("select * from test")
	if err != nil {
		println(err.Error())
		return
	}
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			println(err.Error())
			return
		}
		println(id, name)
	}

}
