package dto

import (
	"crypto_anomaly_searcher/api/constants"
)

type TickerData struct {
	Symbol string `json:"symbol"`
	// PriceChangePercent string `json:"priceChangePercent"`
	Volume15m string `json:"volume"`
	Volume2h  string `json:"volume"`
	Volume1d  string `json:"volume"`
}

type TickerDataList []TickerData

type TickerWinSizeVol struct {
	Ticker     string
	WindowSize constants.WindowSize
	Volume     string
}

func (resultDataList *TickerDataList) SetTickerData(data TickerWinSizeVol) *TickerDataList {
	exist := false
	for _, resTickerData := range *resultDataList {
		if resTickerData.Symbol == data.Ticker {
			exist = true
		}
	}
	if !exist {
		*resultDataList = append(*resultDataList, TickerData{
			Symbol:    data.Ticker,
			Volume15m: "",
			Volume2h:  "",
			Volume1d:  "",
		})
	}
	for i, resTickerData := range *resultDataList {
		if resTickerData.Symbol == data.Ticker {
			switch data.WindowSize {
			case constants.M15:
				resTickerData.Volume15m = data.Volume
			case constants.H2:
				resTickerData.Volume2h = data.Volume
			case constants.D1:
				resTickerData.Volume1d = data.Volume
			}
			(*resultDataList)[i] = resTickerData
			break
		}
	}
	return resultDataList
}
