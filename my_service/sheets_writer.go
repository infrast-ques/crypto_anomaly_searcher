package my_service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"crypto_anomaly_searcher/api"
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
	Ticker      string
	Volume      string
	PriceChange string
}

func ClearSheet() {
	_, err := srv.Spreadsheets.Values.Clear(ssId, "A1:C10000", &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		logrus.Errorf("Unable to clear data table: %v", err)
	}
}

func WriteData(data []sheetTemplate) {
	var values [][]interface{}
	values = append(values, []interface{}{"Ticker", "Volume", "Price Change"})

	for _, datum := range data {
		values = append(values, []interface{}{datum.Ticker, datum.Volume, datum.PriceChange})
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

func ToSheetData(t api.TickerRespList) []sheetTemplate {
	var res []sheetTemplate
	for _, tData := range t {
		volume, errV := strconv.ParseFloat(tData.Volume, 64)
		if errV != nil {
			logrus.Error("Parse ticker volume value error: ", errV)
		}
		if volume == 0.0 {
			continue
		}

		priceChange, errPr := strconv.ParseFloat(tData.PriceChangePercent, 64)
		if errPr != nil {
			logrus.Error("Parse ticker price change value error: ", errPr)
		}
		if priceChange == 0.0 {
			continue
		}

		template := sheetTemplate{
			Ticker:      tData.Symbol,
			Volume:      strings.Replace(fmt.Sprintf("%.2f", volume), ".", ",", 1),
			PriceChange: strings.Replace(fmt.Sprintf("%.2f", priceChange), ".", ",", 1),
		}

		res = append(res, template)
	}

	return res
}
