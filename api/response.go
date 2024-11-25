package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ResponseWrapper struct {
	*http.Response
}

func (r ResponseWrapper) Deserialize(model interface{}) (interface{}, error) {

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Warn("The response reading stream has not closed")
		}
	}(r.Body)

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
