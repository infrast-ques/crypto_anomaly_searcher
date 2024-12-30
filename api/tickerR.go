package api

import (
	"net/http"
	"net/url"
	"strconv"

	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/service"
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

type tickerResp struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
}

type TickerResp struct {
	Ticker          string
	PrChange        float64
	PrChangePercent float64
	Vol             float64
	QuoteVol        float64
}
type TickerRespList []TickerResp

type TickersToWSizeResp struct {
	TickerRespList TickerRespList
	WSize          constants.WindowSize
}

func GetTickersData(tickers []string, period constants.WindowSize) TickerRespList {
	request := tickerRequest(tickers, period)
	response := clientBinance.Send(&request)
	stringModel := utils.Deserialize(response, []tickerResp{})
	res := TickerRespList{}
	for _, data := range stringModel {
		priceChange := stringToFloat(data.PriceChange)
		priceChangePercent := stringToFloat(data.PriceChangePercent)

		volume := stringToFloat(data.Volume)
		if volume == 0.0 {
			service.Logger.Infof("Ticker %s with zero volume", data.Symbol)
			continue
		}

		quoteVolume := stringToFloat(data.QuoteVolume)
		if quoteVolume == 0.0 {
			service.Logger.Infof("Ticker %s with zero quoteVolume", data.Symbol)
			continue
		}

		res = append(res, TickerResp{
			Ticker:          data.Symbol,
			PrChange:        priceChange,
			PrChangePercent: priceChangePercent,
			Vol:             volume,
			QuoteVol:        quoteVolume,
		})
	}
	return res
}

func stringToFloat(v string) float64 {
	res, err := strconv.ParseFloat(v, 64)
	if err != nil {
		service.Logger.Error(err)
	}
	return res
}
