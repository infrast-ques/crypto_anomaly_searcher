package api

import (
	"net/http"
	"net/url"
	"strconv"

	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
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

type tickerResp struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
}

type TickerResp struct {
	Ticker             string
	PriceChange        float64
	PriceChangePercent float64
	Volume             float64
	QuoteVolume        float64
}

type TickersToWindowSizeResp struct {
	TickerRespList TickerRespList
	WindowSize     constants.WindowSize
}

type TickerRespList []TickerResp

func GetTickersData(tickers []string, period constants.WindowSize) TickerRespList {
	request := tickerRequest(tickers, period)
	response := clientBinance.Send(&request)
	stringModel := utils.Deserialize(response, []tickerResp{})
	res := TickerRespList{}
	for _, data := range stringModel {
		priceChange, errPC := strconv.ParseFloat(data.PriceChange, 64)
		if errPC != nil {
			logrus.Error(errPC)
		}
		priceChangePercent, errPCP := strconv.ParseFloat(data.PriceChangePercent, 64)
		if errPCP != nil {
			logrus.Error(errPCP)
		}
		volume, errV := strconv.ParseFloat(data.Volume, 64)
		if errV != nil {
			logrus.Error(errV)
		}
		if volume == 0.0 {
			logrus.Infof("Ticker %s with zero volume", data.Symbol)
		}
		quoteVolume, errQV := strconv.ParseFloat(data.QuoteVolume, 64)
		if errQV != nil {
			logrus.Error(errQV)
		}
		if quoteVolume == 0.0 {
			logrus.Infof("Ticker %s with zero quoteVolume", data.Symbol)
		}
		res = append(res, TickerResp{
			Ticker:             data.Symbol,
			PriceChange:        priceChange,
			PriceChangePercent: priceChangePercent,
			Volume:             volume,
			QuoteVolume:        quoteVolume,
		})
	}
	return res
}
