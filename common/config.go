package common

import (
	"os"

	"github.com/sirupsen/logrus"
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
	UpdateTime float32  `yaml:"update-time-min"`
	SsIds      []string `yaml:"ss-ids"`
}

func initConfig() ConfigM {
	file, err := os.Open("config.yml")
	if err != nil {
		logrus.Fatalf("Open file with error: %v", err)
	}
	defer file.Close()

	var config ConfigM
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		logrus.Fatalf("Decode error YAML: %v", err)
	}
	file.Close()
	return config
}
