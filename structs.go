package main

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type dog struct {
	id   int
	name string
	age  int
}
