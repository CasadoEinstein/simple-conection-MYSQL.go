package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test table")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO testable2 VALUES(23)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("yay, values added")
}
