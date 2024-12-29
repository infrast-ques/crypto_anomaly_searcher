package data_handler

import (
	"fmt"

	"crypto_anomaly_searcher/api"
	"crypto_anomaly_searcher/service/data_collector/dto"
)

func StrFmtComputedData(
	data dto.TickerRawDataList,
	header string,
	computer func(dto.TickerRawData) string,
) []string {
	res := []string{header}
	for _, d := range data {
		res = append(res, computer(d))
	}
	for _, re := range res {
		fmt.Println(re)
	}
	return res
}

func MergeTickersData(tickerDataLists []api.TickersToWSizeResp) dto.TickerRawDataList {
	var resultTickerDataList []dto.TickerVolByWSize
	for _, tickerDataList := range tickerDataLists {
		for _, tickerData := range tickerDataList.TickerRespList {
			hui := dto.TickerVolByWSize{
				Ticker:      tickerData.Ticker,
				WSize:       tickerDataList.WSize,
				Vol:         tickerData.Vol,
				PrChngPrcnt: tickerData.PrChangePercent,
			}
			resultTickerDataList = append(resultTickerDataList, hui)
		}
	}
	var mergeResult dto.TickerRawDataList
	for _, v := range resultTickerDataList {
		mergeResult.SetTickerData(v)
	}
	return mergeResult
}
