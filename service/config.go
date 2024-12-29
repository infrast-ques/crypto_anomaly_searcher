package service

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigData = initConfig()

type ConfigM struct {
	Telegram telegramConfig `yaml:"telegram"`
	Sheet    sheetConfig    `yaml:"sheet"`
}

type telegramConfig struct {
	ApiKey     string `yaml:"api-key"`
	TestUserId int64  `yaml:"test-user-id"`
}

type sheetConfig struct {
	Raw       sheetData `yaml:"raw-data"`
	Computed1 sheetData `yaml:"computed1"`
}

type sheetData struct {
	SsIds []string `yaml:"ss-ids"`
}

func initConfig() ConfigM {
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Open file with error: %v", err)
	}
	defer file.Close()

	var config ConfigM
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Decode error YAML: %v", err)
	}
	file.Close()
	return config
}
