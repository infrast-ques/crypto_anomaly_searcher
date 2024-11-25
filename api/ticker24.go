package api

import (
	"net/http"
	"net/url"
)

const (
	endpoint    = "/api/v3/ticker/24hr"
	binanceHost = "api.binance.com"
)

var tickers = []string{
	"BTCUSDT",
}
var params = map[string]string{
	"symbol": tickers[0],
}

var request = http.Request{
	Method: "GET",
	URL: &url.URL{
		Scheme:   "https",
		Host:     binanceHost,
		Path:     endpoint,
		RawQuery: MapToQueryParams(params),
	},
	Close: true,
}

var Ticker24Req = Request[TickerResp]{
	Req:   request,
	Model: TickerResp{},
}

type TickerResp struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`  // Используем int64 для времени
	CloseTime          int64  `json:"closeTime"` // Используем int64 для времени
	FirstID            int    `json:"firstId"`   // Первичный идентификатор
	LastID             int    `json:"lastId"`    // Последний идентификатор
	Count              int    `json:"count"`     // Количество сделок
}
