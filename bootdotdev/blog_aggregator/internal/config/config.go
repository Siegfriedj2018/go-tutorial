package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"
// This has to be changed on other os
const configFilePath = "/mnt/c/Users/siegf/SynologyDrive/Drive/Go_Tutorial/bootdotdev/blog_aggregator"

type Config struct {
	DbUrl 				string  `json:"db_url"`
	CurrentUser		string  `json:"current_user_name"`
}


func Read() Config {
	path, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("there was an invalid filepath: %v", err)
	}
	
	var configData Config
	data, err := os.Open(path)
	if err != nil {
		log.Fatalf("there was an error opening the file: %v", err)
	}
	defer data.Close()
	

	decoder := json.NewDecoder(data)
	err = decoder.Decode(&configData)
	if err != nil {
		log.Fatalf("there was an error unmarshalling the data: %v", err)
	}
	
	return configData
}

func getConfigFilePath() (string, error) {
	fullPath := filepath.Join(configFilePath, configFileName)
	return fullPath, nil
}

func write(conf *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("there was an error getting config path: %w", err)
	}

	data, err := os.Create(path)
	userData := json.NewEncoder(data)
	err = userData.Encode(conf)
	if err != nil {
		return fmt.Errorf("there was an error encoding the data: %w", err)
	}

	return nil
}

func (confData *Config) SetUser(username string) error {
	confData.CurrentUser = username
	return write(confData)
}