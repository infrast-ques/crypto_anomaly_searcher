package data_collector

import (
	"errors"
	"strings"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/service"
	"crypto_anomaly_searcher/service/data_collector/dto"
	"crypto_anomaly_searcher/service/data_handler"
	"crypto_anomaly_searcher/utils"
)

type batchTickers [][]string

type AggregatedData struct {
	RawData       dto.TickerRawDataList
	Computed1Data []string
}

func AggregateData() AggregatedData {
	tickers, _ := getTickerList()
	tickerDataLists := collectTickersData(tickers)
	rawdata := data_handler.MergeTickersData(tickerDataLists)
	computed1Data := data_handler.StrFmtComputedData(
		rawdata,
		data_handler.Compute2ListHeader,
		data_handler.Compute2List,
	)
	return AggregatedData{
		RawData:       rawdata,
		Computed1Data: computed1Data,
	}
}

func getTickerList() (batchTickers, error) {
	allTickers := api.GetAllTickers()
	if len(allTickers) == 0 {
		error := errors.New("zero tickers in response")
		service.Logger.Error(error)
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

func collectTickersData(batchTickers batchTickers) []api.TickersToWSizeResp {
	var tickersRespLists []api.TickersToWSizeResp
	// todo тут можно горутины воткнуть
	for _, size := range []constants.WindowSize{constants.M15, constants.H2, constants.D1} {
		tickersRespLists = append(
			tickersRespLists,
			api.TickersToWSizeResp{
				TickerRespList: getTickersData(batchTickers, size),
				WSize:          size,
			},
		)
	}
	return tickersRespLists
}

func getTickersData(batchTickers batchTickers, period constants.WindowSize) api.TickerRespList {
	flatTickerDataList := make(api.TickerRespList, 0)
	for _, tickers := range batchTickers {
		// todo воткнуть горутины
		tickersData := api.GetTickersData(tickers, period)
		if len(tickersData) != len(tickers) {
			service.Logger.Warnf("Error: received an incorrect number of tickers. Expected: %d, Received: %d", len(tickers), len(tickersData))
		}
		for _, tickerData := range tickersData {
			flatTickerDataList = append(flatTickerDataList, tickerData)
		}
	}
	return flatTickerDataList
}
