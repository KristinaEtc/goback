package main

import (
	"./connection"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB

type Structure struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

func PrepareRequest(data interface{}, err error) *Structure {
	resp := new(Structure)
	resp.Data = data
	resp.Error = ""
	resp.Status = http.StatusOK

	if err != nil {
		resp.Error = err.Error()
		resp.Data = nil
		resp.Status = http.StatusNotFound
	}

	return resp
}

func main() {
	db = connection.CreateConnection()
	router := mux.NewRouter()
	router.HandleFunc("/games", GetGamesEndpoint).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}", GetGameEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}
