package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Serialize(model interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(model, "", "	")
	if err != nil {
		logrus.Error(errors.New("Serialization error: " + err.Error()))
		return "", err
	}
	return string(jsonData), nil
}

func Deserialize[T any](r *http.Response, model T) T {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Warn("The response reading stream has not closed")
		}
	}(r.Body)

	responseBytes, err := io.ReadAll(r.Body)
	logrus.Info("Response body - " + string(responseBytes))

	if err != nil {
		logrus.Error(errors.New("ReadAll Body - " + err.Error()))
	}

	err = json.Unmarshal(responseBytes, &model)
	if err != nil {
		logrus.Error(errors.New("json.Unmarshal - " + err.Error()))
	}

	return model
}
