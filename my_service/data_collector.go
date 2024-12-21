package my_service

import (
	"strings"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/my_service/dto"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
)

func AggregateData() dto.TickerDataList {
	tickersParts := getTickerList()
	tickersRespLists := collectTickerData(tickersParts)
	res := mergeTickersData(tickersRespLists)
	return res
}

func collectTickerData(tickersParts [][]string) []api.TickersToWindowSizeResp {
	var tickersRespLists []api.TickersToWindowSizeResp
	// todo тут можно горутины воткнуть
	for _, size := range []constants.WindowSize{constants.M15, constants.H2, constants.D1} {
		tickersRespLists = append(
			tickersRespLists,
			api.TickersToWindowSizeResp{
				TickerRespList: getTickersData(tickersParts, size),
				WindowSize:     size,
			},
		)
	}
	return tickersRespLists
}

func mergeTickersData(tickerDataLists []api.TickersToWindowSizeResp) dto.TickerDataList {
	var resultTickerDataList []dto.TickerWinSizeVol
	for _, tickerDataList := range tickerDataLists {
		for _, tickerData := range tickerDataList.TickerRespList {
			hui := dto.TickerWinSizeVol{
				Ticker:     tickerData.Symbol,
				WindowSize: tickerDataList.WindowSize,
				Volume:     tickerData.Volume,
			}
			resultTickerDataList = append(resultTickerDataList, hui)
		}
	}
	var mergeResult dto.TickerDataList
	for _, v := range resultTickerDataList {
		mergeResult.SetTickerData(v)
	}
	return mergeResult
}

func getTickerList() [][]string {
	tickers := api.GetAllTickers()
	filteredTickers := tickersFilter(tickers)[0:2]
	var tickersParts [][]string
	for start := 0; start < len(filteredTickers); start += 100 {
		end := start + 100
		if end > len(filteredTickers) {
			end = len(filteredTickers)
		}
		tickersParts = append(tickersParts, filteredTickers[start:end])
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
