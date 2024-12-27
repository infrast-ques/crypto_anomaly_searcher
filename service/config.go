package service

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigData = initConfig()

type ConfigModel struct {
	Telegram telegramConfig `yaml:"telegram"`
}

type telegramConfig struct {
	ApiKey     string `yaml:"api-key"`
	TestUserId int64  `yaml:"test-user-id"`
}

func initConfig() ConfigModel {
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Open file with error: %v", err)
	}
	defer file.Close()

	var config ConfigModel
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Decode error YAML: %v", err)
	}
	file.Close()
	return config
}
