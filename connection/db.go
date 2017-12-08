package connection

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Config struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Cannot find config file:", err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func CreateConnection() *sql.DB {
	var config Config
	var db *sql.DB
	config = loadConfiguration("../goback/connection/config.json")
	var err error
	db, err = sql.Open(
		config.Driver,
		config.User+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Name)
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
	db.SetMaxOpenConns(10)

	return db
}
