package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	db_host = "localhost"
	db_port = 3306
	db_user = "root"
	db_pass = "root"
	db_name = "admin"
	db_type = "mysql"
)

func main() {
	var sqlstring string
	sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8")
	//db, err := sql.Open(db_type, dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}
