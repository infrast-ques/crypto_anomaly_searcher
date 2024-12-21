package my_service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"crypto_anomaly_searcher/my_service/dto"
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
	Ticker    string
	Volume15m string
	Volume2h  string
	Volume1d  string
	// PriceChange string
}

func ClearSheet() {
	_, err := srv.Spreadsheets.Values.Clear(ssId, "A1:C10000", &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		logrus.Errorf("Unable to clear data table: %v", err)
	}
}

func WriteData(data []sheetTemplate) {
	var values [][]interface{}
	values = append(values, []interface{}{"Ticker", "Volume15m", "Volume2h", "Volume1d"})

	for _, datum := range data {
		values = append(values, []interface{}{
			datum.Ticker,
			datum.Volume15m,
			datum.Volume2h,
			datum.Volume1d,
			// datum.PriceChange,
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
		volume15m, errV := strconv.ParseFloat(tData.Volume15m, 64)
		if errV != nil {
			logrus.Error("Parse ticker volume value error: ", errV)
		}
		if volume15m == 0.0 {
			continue
		}
		volume2h, errV := strconv.ParseFloat(tData.Volume2h, 64)
		if errV != nil {
			logrus.Error("Parse ticker volume value error: ", errV)
		}
		if volume2h == 0.0 {
			continue
		}
		volume1d, errV := strconv.ParseFloat(tData.Volume1d, 64)
		if errV != nil {
			logrus.Error("Parse ticker volume value error: ", errV)
		}
		if volume1d == 0.0 {
			continue
		}

		// priceChange, errPr := strconv.ParseFloat(tData.PriceChangePercent, 64)
		// if errPr != nil {
		// 	logrus.Error("Parse ticker price change value error: ", errPr)
		// }
		// if priceChange == 0.0 {
		// 	continue
		// }

		template := sheetTemplate{
			Ticker:    tData.Symbol,
			Volume15m: strings.Replace(fmt.Sprintf("%.2f", volume15m), ".", ",", 1),
			Volume2h:  strings.Replace(fmt.Sprintf("%.2f", volume2h), ".", ",", 1),
			Volume1d:  strings.Replace(fmt.Sprintf("%.2f", volume1d), ".", ",", 1),
			// PriceChange: strings.Replace(fmt.Sprintf("%.2f", priceChange), ".", ",", 1),
		}

		res = append(res, template)
	}

	return res
}
