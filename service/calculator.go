package service

import (
	"fmt"

	"crypto_anomaly_searcher/service/dto"
	"crypto_anomaly_searcher/utils"
)

func compute(data dto.TickerData) string {
	avgVolFirst22h := (data.Volume1d - data.Volume2h) / 11
	avgVolFirst45m := (data.Volume2h - data.Volume15m) / 3
	last2hVolPercentFromAvg22hVol := data.Volume2h/avgVolFirst22h*100 - 100
	last15mVolPercentFromAvg2hVol := data.Volume15m/avgVolFirst45m*100 - 100
	resStr := fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s\t%s",
		data.Ticker,
		utils.FloatToStrFmt(last15mVolPercentFromAvg2hVol),
		utils.FloatToStrFmt(last2hVolPercentFromAvg22hVol),
		utils.FloatToStrFmt(data.PriceChangePercent15m),
		utils.FloatToStrFmt(data.PriceChangePercent15m),
		utils.FloatToStrFmt(data.PriceChangePercent1d),
	)
	return resStr
}
