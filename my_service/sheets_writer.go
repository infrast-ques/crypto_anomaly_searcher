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

func WriteData(data []sheetTemplate) {
	srv := getSheets()
	var valueRange sheets.ValueRange
	valueRange.Values = append(valueRange.Values, []interface{}{"Ticker", "Volume", "Price Change"})
	srv.Spreadsheets.Values.Update(ssId, "A1", &valueRange).ValueInputOption("RAW").Do()

	for i, datum := range data {
		var valueRange sheets.ValueRange
		valueRange.Values = append(valueRange.Values, []interface{}{datum.Ticker, datum.Volume, datum.PriceChange})
		srv.Spreadsheets.Values.
			Update(ssId, fmt.Sprintf("A%d", i+2), &valueRange).
			ValueInputOption("USER_ENTERED").
			Do()
	}
}

func ToSheetData(t api.TickerPespList) []sheetTemplate {
	var res []sheetTemplate
	for _, tData := range t {
		volume, err := strconv.ParseFloat(tData.Volume, 64)
		if err != nil {
			logrus.Error("Parse ticker volume value error: ", err)
		}
		priceChange, err := strconv.ParseFloat(tData.PriceChangePercent, 64)
		if err != nil {
			logrus.Error("Parse ticker price change value error: ", err)
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
