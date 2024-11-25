package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"crypto_anomaly_searcher/api"
	"github.com/sirupsen/logrus"
)

var q = http.Request{}

func main() {

	resp, err := ApiClientBinance.Send(&api.Ticker24Req)

	if err != nil {
		logrus.Error("Request error")
	}
	tickerResp := api.TickerResp{}
	api.Deserialize(resp, &tickerResp)
	json, _ := json.MarshalIndent(tickerResp, "", "")
	fmt.Println(string(json))
}
