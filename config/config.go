package config

import (
	"encoding/json"
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

const configPath = "../goback/config.json"

func GetConfig() Config {
	var config Config
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Cannot find config file:", err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
