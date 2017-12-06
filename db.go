package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/gameschanger")
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
}
