package api

import (
	"errors"
	"fmt"
	"net/http"

	"crypto_anomaly_searcher/service"
)

type client struct {
	client  *http.Client
	headers http.Header
}

var clientBinance = client{
	client:  http.DefaultClient,
	headers: nil,
}

func (a client) Send(r *http.Request) *http.Response {
	for key, values := range r.Header {
		for _, value := range values {
			a.headers.Set(key, value)
		}
	}

	service.Logger.Info(fmt.Sprintf("Request - %s: %s \n%s", r.Method, r.URL, r.Body))

	resp, err := a.client.Do(r)
	// todo add request utils.Logger

	if err != nil {
		service.Logger.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		service.Logger.Warn(errors.New("Status code:" + resp.Status))
	}

	return resp
}
