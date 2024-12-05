package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

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
	// todo logrus.Info("Response body - " + string(responseBytes))

	if err != nil {
		logrus.Error(errors.New("ReadAll Body - " + err.Error()))
	}

	err = json.Unmarshal(responseBytes, &model)
	if err != nil {
		logrus.Error(errors.New("json.Unmarshal - " + err.Error()))
	}

	return model
}

func MapToQueryParams(paramsMap map[string]string) string {
	if len(paramsMap) == 0 {
		return ""
	}

	var sBuilder = strings.Builder{}
	for p, pVal := range paramsMap {
		sBuilder.WriteString(p + "=" + pVal + "&")
	}

	return sBuilder.String()
}

func AsQueryParamList(list []string) string {
	return "[\"" + strings.Join(list, "\",\"") + "\"]"
}

// Flatten todo допилить
func Flatten[T [][]interface{}](lists T) []interface{} {
	var res []interface{}
	for _, list := range lists {
		for _, item := range list {
			res = append(res, item)
		}
	}
	return res
}

func SliceFilter(tickers []string, predicate func(string) bool) []string {
	var res []string
	for _, ticker := range tickers {
		if predicate(ticker) {
			res = append(res, ticker)
		}
	}
	return res
}
