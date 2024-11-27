package api

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
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

	resp, err := a.client.Do(r)
	// todo add request log

	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Warn(errors.New("Status code:" + resp.Status))
	}

	return resp
}
