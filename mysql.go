package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	var result sql.Result
	result, err = db.Exec("insert into user(name, age) values(?,?)", "test", 14)
	if err != nil {
		fmt.Println(err)
		return
	}

	lastId, _ := result.LastInsertId()
	fmt.Println("新插入记录的ID为", lastId)

	var row *sql.Row
	row = db.QueryRow("select * from user")
	var name string
	var id, age int
	err = row.Scan(&id, &name, &age)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id, "\t", name, "\t", age)

	result, err = db.Exec("insert into user(name, age) values(?,?)", "eee", 23)

	var rows *sql.Rows
	rows, err = db.Query("select * from user")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var name string
		var id, age int
		rows.Scan(&id, &age, &name)
		fmt.Println(id, "\t", name, "\t", age)
	}
	rows.Close()

	//db.Exec("truncate table user")
}
