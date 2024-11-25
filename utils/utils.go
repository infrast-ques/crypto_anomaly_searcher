package utils

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func Serialize(model interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(model, "", "")
	if err != nil {
		logrus.Error("Serialization error: " + err.Error())
		return "", err
	}
	return string(jsonData), nil
}
