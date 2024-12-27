package service

import (
	"context"

	"crypto_anomaly_searcher/service/dto"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	jsonKeyPath = "google_sheets_key.json"
	ssId        = "1z4eXjtP9WwKpyh4I_FZbBue3_Az2WPo18XZnGMTQsRk"
	srv         = getSheets()
)

func getSheets() *sheets.Service {
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsFile(jsonKeyPath))
	if err != nil {
		logrus.Error(err)
	}
	return srv
}

type sheetTemplate struct {
	Ticker        string
	Volume15m     string
	Volume2h      string
	Volume1d      string
	PriceChange2h string
}

func ClearSheet() {
	_, err := srv.Spreadsheets.Values.Clear(ssId, "A1:N10000", &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		logrus.Errorf("Unable to clear data table: %v", err)
	}
}

func WriteData(t dto.TickerDataList) {
	data := ToSheetData(t)
	var values [][]interface{}
	values = append(values, []interface{}{
		"Ticker",
		"Volume15m",
		"Volume2h",
		"Volume1d",
		"PriceChange2h",
	})

	for _, datum := range data {
		values = append(values, []interface{}{
			datum.Ticker,
			datum.Volume15m,
			datum.Volume2h,
			datum.Volume1d,
			datum.PriceChange2h,
		})
	}
	valueRange := sheets.ValueRange{
		Values: values,
	}

	_, err := srv.Spreadsheets.Values.
		Update(ssId, "A1", &valueRange).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		logrus.Errorf("Unable to write data to sheet: %v", err)
	}
}

func ToSheetData(t dto.TickerDataList) []sheetTemplate {
	var res []sheetTemplate
	for _, tData := range t {
		template := sheetTemplate{
			Ticker:        tData.Ticker,
			Volume15m:     utils.FloatToStrFmt(tData.Volume15m),
			Volume2h:      utils.FloatToStrFmt(tData.Volume2h),
			Volume1d:      utils.FloatToStrFmt(tData.Volume1d),
			PriceChange2h: utils.FloatToStrFmt(tData.PriceChangePercent2h),
		}
		res = append(res, template)
	}

	return res
}
