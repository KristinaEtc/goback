package main

import "log"

func getUsers() []*user {
	users, err := db.Query("SELECT id,username,email FROM user")

	if err != nil {
		log.Fatal(err)
	}

	result := make([]*user, 0)
	for users.Next() {
		usr := new(user)
		err := users.Scan(&usr.id, &usr.username, &usr.email)
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
