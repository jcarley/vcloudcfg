package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Token    string `json:token`
	Username string `json:username`
}

func NewConfig() (*Config, error) {

	file, err := os.Open("config.json")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("config.json does not exist", err)
		}
		log.Fatalln("Failed to open config.json", err)
	}
	defer file.Close()

	var config *Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
