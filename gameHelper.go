package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Year        string `json:"year"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

func GetGames() *Structure {
	games, err := db.Query("SELECT id,name,year,cover,description FROM game LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	result := make([]*game, 0)
	for games.Next() {
		game := new(game)
		err := games.Scan(
			&game.Id,
			&game.Name,
			&game.Year,
			&game.Cover,
			&game.Description)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, game)
	}
	resp := PrepareRequest(result, err)
	return resp
}
func GetGame(id string) *Structure {
	fmt.Println(id)
	row := db.QueryRow("SELECT id,name,year,cover,description FROM game WHERE id = ?", id)
	game := new(game)
	err := row.Scan(
		&game.Id,
		&game.Name,
		&game.Year,
		&game.Cover,
		&game.Description)
	resp := PrepareRequest(game, err)
	return resp
}

func GetGamesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(GetGames())
}
func GetGameEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(GetGame(mux.Vars(req)["id"]))
}
