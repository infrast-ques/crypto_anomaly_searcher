package my_service

import (
	"crypto_anomaly_searcher/api"
)

func GetCryptoData() api.TickerPespList {
	// todo Debug
	// tickers := api.GetAllTickers()[0:2]
	// tickersData := api.GetTickersData(tickers)
	tickersData := api.GetTickersData([]string{"BTCUSDT"})
	return tickersData
}
