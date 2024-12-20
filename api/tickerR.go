package api

import (
	"net/http"
	"net/url"

	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/utils"
)

func tickerRequest(tickers []string, windowSize constants.WindowSize) http.Request {
	return http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   binanceHost,
			Path:   tickersEndpoint,
			RawQuery: utils.MapToQueryParams(map[string]string{
				"symbols":    utils.AsQueryParamList(tickers),
				"windowSize": string(windowSize),
			}),
		},
		Close: true,
	}
}

type TickerResp struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
}

type TickerRespList []TickerResp

func GetTickersData(tickers []string, period constants.WindowSize) TickerRespList {
	request := tickerRequest(tickers, period)
	response := clientBinance.Send(&request)
	return utils.Deserialize(response, TickerRespList{})
}
