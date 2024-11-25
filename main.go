package main

import (
	"fmt"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
)

func main() {

	resp, err := api.ClientBinance.Send(&api.Ticker24Req.Req)

	if err != nil {
		logrus.Error("Request error")
	}
	tickerResp := api.TickerResp{}
	resp.Deserialize(&tickerResp) // перенести десериализацию сразу в send
	json, _ := utils.Serialize(tickerResp)
	fmt.Println(json)
}
