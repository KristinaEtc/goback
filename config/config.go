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

var config Config

func GetConfig() Config {
	configFile := readFile("../goback/config.json")
	createConfig(configFile, "json")

	return config
}

func readFile(path string) *os.File {
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Cannot find config file:", err)
	}
	return configFile
}

func createConfig(file *os.File, format string) {
	switch format {
	case "json":
		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&config)
	default:
		log.Fatal("Config file type is not supported")
	}
}
