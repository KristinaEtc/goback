package main

import (
	"./connection"
	"database/sql"
	"log"
)

func getUsers() []*user {
	var db *sql.DB
	db = connection.CreateConnection()
	users, err := db.Query("SELECT id,username,email FROM user")
	if err != nil {
		log.Fatal(err)
	}
	result := make([]*user, 0)
	for users.Next() {
		usr := new(user)
		err := users.Scan(&usr.Id, &usr.Username, &usr.Email)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, usr)
	}
	if err = users.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
