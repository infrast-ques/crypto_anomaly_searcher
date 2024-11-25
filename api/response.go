package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Deserialize(r *http.Response, model interface{}) (interface{}, error) {
	defer r.Body.Close()
	responseBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Error reading response body:", err)
		return nil, err
	}
	err = json.Unmarshal(responseBytes, &model)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	return model, nil
}
