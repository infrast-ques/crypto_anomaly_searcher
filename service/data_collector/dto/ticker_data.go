package dto

import (
	"reflect"
	"strings"

	"crypto_anomaly_searcher/api/constants"
	"crypto_anomaly_searcher/service"
	"crypto_anomaly_searcher/utils"
)

type TickerRawData struct {
	Ticker          string
	Vol15m          float64
	Vol2h           float64
	Vol1d           float64
	PrcChngPrcnt15m float64
	PrcChngPrcnt2h  float64
	PrcChngPrcnt1d  float64
}

func (t TickerRawData) toString() string {
	return t.Ticker + service.SheetSplitter +
		utils.FloatToStrFmt(t.Vol15m) + service.SheetSplitter +
		utils.FloatToStrFmt(t.Vol2h) + service.SheetSplitter +
		utils.FloatToStrFmt(t.Vol1d) + service.SheetSplitter +
		utils.FloatToStrFmt(t.PrcChngPrcnt15m) + service.SheetSplitter +
		utils.FloatToStrFmt(t.PrcChngPrcnt2h) + service.SheetSplitter +
		utils.FloatToStrFmt(t.PrcChngPrcnt1d)
}

type TickerRawDataList []TickerRawData

type TickerVolByWSize struct {
	Ticker      string
	WSize       constants.WindowSize
	Vol         float64
	PrChngPrcnt float64
}

func (ts *TickerRawDataList) ToStringList() []string {
	r := reflect.TypeOf(TickerRawData{})
	var columnNames []string
	for i := 0; i < r.NumField(); i++ {
		columnNames = append(columnNames, r.Field(i).Name)
	}
	result := []string{strings.Join(columnNames, service.SheetSplitter)}

	for _, data := range *ts {
		result = append(result, data.toString())
	}
	return result
}

func (ts *TickerRawDataList) SetTickerData(data TickerVolByWSize) *TickerRawDataList {
	exist := false
	for _, resTickerData := range *ts {
		if resTickerData.Ticker == data.Ticker {
			exist = true
		}
	}
	if !exist {
		*ts = append(*ts, TickerRawData{
			Ticker:          data.Ticker,
			Vol15m:          0,
			Vol2h:           0,
			Vol1d:           0,
			PrcChngPrcnt15m: 0,
			PrcChngPrcnt2h:  0,
			PrcChngPrcnt1d:  0,
		})
	}
	// todo подумать как переделать на мапу
	for i, resTickerData := range *ts {
		if resTickerData.Ticker == data.Ticker {
			switch data.WSize {
			case constants.M15:
				resTickerData.Vol15m = data.Vol
				resTickerData.PrcChngPrcnt15m = data.PrChngPrcnt
			case constants.H2:
				resTickerData.Vol2h = data.Vol
				resTickerData.PrcChngPrcnt2h = data.PrChngPrcnt
			case constants.D1:
				resTickerData.Vol1d = data.Vol
				resTickerData.PrcChngPrcnt1d = data.PrChngPrcnt
			}
			(*ts)[i] = resTickerData
			break
		}
	}
	return ts
}
