package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitSql(sourceName string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", sourceName)
	return db, err
}

type Mem struct {
	Id int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age int `json:"age" db:"age"`
	Email string `json:"email" db:"email"`
}


// 在查询的数据库这里需要封装错误返回
func retrieveStudent(name string) (ret Mem, err error) {
	rows, err := db.Query(`select * from member where name=?`, name)
	if err != nil {
		fmt.Errorf("search student %s fail: %w", name, err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&ret)
		if err != nil {
			fmt.Errorf("rows sacn mem fail: %w", err)
		}
	}
	return
}
