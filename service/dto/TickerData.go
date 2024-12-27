package dto

import (
	"crypto_anomaly_searcher/api/constants"
)

type TickerData struct {
	Ticker                string
	Volume15m             float64
	PriceChangePercent15m float64
	Volume2h              float64
	PriceChangePercent2h  float64
	Volume1d              float64
	PriceChangePercent1d  float64
}

type TickerDataList []TickerData

type TickerWinSizeVol struct {
	Ticker             string
	WindowSize         constants.WindowSize
	Volume             float64
	PriceChangePercent float64
}

func (resultDataList *TickerDataList) SetTickerData(data TickerWinSizeVol) *TickerDataList {
	exist := false
	for _, resTickerData := range *resultDataList {
		if resTickerData.Ticker == data.Ticker {
			exist = true
		}
	}
	if !exist {
		*resultDataList = append(*resultDataList, TickerData{
			Ticker:               data.Ticker,
			PriceChangePercent2h: 0,
			Volume15m:            0,
			Volume2h:             0,
			Volume1d:             0,
		})
	}
	for i, resTickerData := range *resultDataList {
		if resTickerData.Ticker == data.Ticker {
			switch data.WindowSize {
			case constants.M15:
				resTickerData.Volume15m = data.Volume
				resTickerData.PriceChangePercent15m = data.PriceChangePercent
			case constants.H2:
				resTickerData.Volume2h = data.Volume
				resTickerData.PriceChangePercent2h = data.PriceChangePercent
			case constants.D1:
				resTickerData.Volume1d = data.Volume
				resTickerData.PriceChangePercent1d = data.PriceChangePercent
			}
			(*resultDataList)[i] = resTickerData
			break
		}
	}
	return resultDataList
}
