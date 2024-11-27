package api

import (
	"net/http"
	"net/url"

	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/utils"
)

func getTickerRequest(tickers []string, windowSize constants.WindowSize) http.Request {
	return http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   binanceHost,
			Path:   tickersEndpoint,
			RawQuery: MapToQueryParams(map[string]string{
				"symbols":    AsQueryParamList(tickers),
				"windowSize": string(windowSize),
			}),
		},
		Close: true,
	}
}

type tickerResp struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
}

func GetTickersData(tickers []string) []tickerResp {
	request := getTickerRequest(tickers, constants.M30)
	response := clientBinance.Send(&request)
	return utils.Deserialize(response, []tickerResp{})
}
