package my_service

import (
	"strings"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
)

func AggregateData() api.TickerRespList {
	tickersParts := getTickersParts()
	// todo тут можно горутины воткнуть
	res1h := getTickersData(tickersParts, constants.H1)
	// res2h := getTickersData(tickersParts, constants.H2)
	// res1d := getTickersData(tickersParts, constants.D1)
	// res7d := getTickersData(tickersParts, constants.D7)

	return res1h
}

func compute() {

}

func getTickersParts() [][]string {
	tickers := api.GetAllTickers()
	filteredTickers := tickersFilter(tickers)
	var tickersParts [][]string
	for i := 0; i < len(filteredTickers); i += 100 {
		end := i + 100
		if end > len(filteredTickers) {
			end = len(filteredTickers)
		}
		tickersParts = append(tickersParts, filteredTickers[i:end])
	}
	return tickersParts
}

func tickersFilter(tickers []string) []string {
	res := utils.SliceFilter(tickers, func(s string) bool {
		return strings.HasSuffix(s, "USDT")
	})
	return res
}

func getTickersData(tickersParts [][]string, period constants.WindowSize) api.TickerRespList {
	resTickerData := make(api.TickerRespList, 0)
	for _, tickers := range tickersParts {
		// todo воткнуть горутины
		data := api.GetTickersData(tickers, period)
		if len(data) != len(tickers) {
			logrus.Errorf(
				"Error: received an incorrect number of tickers. Expected: %d, Received: %d",
				len(tickers),
				len(data),
			)
		}
		for _, datum := range data {
			resTickerData = append(resTickerData, datum)
		}
	}
	return resTickerData
}
