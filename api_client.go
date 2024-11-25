package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var ApiClientBinance = ApiClient{
	client:  http.DefaultClient,
	headers: nil,
}

type ApiClient struct {
	client  *http.Client
	headers http.Header
}

func (a ApiClient) Send(r *http.Request) (*http.Response, error) {
	for key, values := range a.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	resp, err := a.client.Do(r)
	if err != nil {
		logrus.Warn(err)
		return resp, err
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Warn("Status code:" + resp.Status)
	}
	return resp, nil
}
