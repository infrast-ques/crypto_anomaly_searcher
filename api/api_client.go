package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var ClientBinance = client{
	client:  http.DefaultClient,
	headers: nil,
}

type client struct {
	client  *http.Client
	headers http.Header
}

func (a client) Send(r *http.Request) (*ResponseWrapper, error) {
	for key, values := range r.Header {
		for _, value := range values {
			a.headers.Set(key, value)
		}
	}

	resp, err := a.client.Do(r)

	// todo add request log

	if err != nil {
		logrus.Warn(err)
		return &ResponseWrapper{Response: resp}, err
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Warn("Status code:" + resp.Status)
	}

	return &ResponseWrapper{Response: resp}, nil
}
