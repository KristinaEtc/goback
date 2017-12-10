package connection

import (
	"../config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CreateConnection() *sql.DB {
	c := config.GetConfig()
	fmt.Println(c)
	db, err := sql.Open(
		c.Driver,
		c.User+":"+c.Password+"@tcp("+c.Host+":"+c.Port+")/"+c.Name)
	if err != nil {
		log.Fatal("Invalid DB config: ", err)
	}
	db.SetMaxOpenConns(10)

	return db
}
