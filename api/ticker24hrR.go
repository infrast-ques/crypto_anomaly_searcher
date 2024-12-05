package api

import (
	"net/http"
	"net/url"

	"crypto_anomaly_searcher/utils"
)

var ticker24hrR = http.Request{
	Method: "GET",
	URL: &url.URL{
		Scheme: "https",
		Host:   binanceHost,
		Path:   tickers24hrEndpoint,
	},
	Close: true,
}

type ticker24hrResp struct {
	Symbol string `json:"symbol"`
	Count  int    `json:"count"`
}

func GetAllTickers() []string {
	response := clientBinance.Send(&ticker24hrR)

	ticker24hr := utils.Deserialize(response, []ticker24hrResp{})

	var tickers []string

	for _, resp := range ticker24hr {
		if resp.Count > 0 {
			tickers = append(tickers, resp.Symbol)
		}
	}

	return tickers
}
