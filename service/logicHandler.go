package service

import (
	"fmt"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/utils"
)

func GetData() {
	tickers := api.GetAllTickers()[0:2]
	tickersData := api.GetTickersData(tickers)
	fmt.Println(utils.Serialize(tickersData))
	fmt.Println(tickersData)
}
