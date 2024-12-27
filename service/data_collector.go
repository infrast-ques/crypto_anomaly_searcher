package service

import (
	"errors"
	"fmt"
	"strings"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/service/dto"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
)

func AggregateData() dto.TickerDataList {
	tickers, _ := getTickerList()
	tickerDataLists := collectTickersData(tickers)
	datas := mergeTickersData(tickerDataLists)
	strFmtComputedData(datas)
	return datas
}

func strFmtComputedData(data dto.TickerDataList) []string {
	res := []string{"ticker\t15m/2h\t2h/24h\t15m%\t2h%\t1d%"}
	for _, d := range data {
		res = append(res, compute(d))
	}
	for _, re := range res {
		fmt.Println(re)
	}
	return res
}

func collectTickersData(tickersParts [][]string) []api.TickersToWindowSizeResp {
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
				Ticker:             tickerData.Ticker,
				WindowSize:         tickerDataList.WindowSize,
				Volume:             tickerData.Volume,
				PriceChangePercent: tickerData.PriceChangePercent,
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

func getTickerList() ([][]string, error) {
	allTickers := api.GetAllTickers()
	if len(allTickers) == 0 {
		error := errors.New("zero tickers in response")
		logrus.Error(error)
		return nil, error
	}
	filteredTickers := tickersFilter(allTickers) // [0:2] // todo для дебага, чтобы не запрашивать данные по всем инструментам
	var tickers [][]string
	for start := 0; start < len(filteredTickers); start += 100 {
		end := start + 100
		if end > len(filteredTickers) {
			end = len(filteredTickers)
		}
		tickers = append(tickers, filteredTickers[start:end])
	}
	return tickers, nil
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
