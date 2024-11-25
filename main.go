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
		logrus.Error("Request error", err)
	}
	tickerResp := api.TickerResp{}
	// todo перенести десериализацию сразу в send
	resp.Deserialize(&tickerResp)
	json, _ := utils.Serialize(tickerResp)
	fmt.Println(json)
}
