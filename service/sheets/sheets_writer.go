package sheets

import (
	"context"
	"log"
	"strings"

	"crypto_anomaly_searcher/service"
	"crypto_anomaly_searcher/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	jsonKeyPath = "google_sheets_key.json"
	srv         = getSheets()
)

func getSheets() *sheets.Service {
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsFile(jsonKeyPath))
	if err != nil {
		logrus.Error(err)
	}
	return srv
}

func FillSheet(ssIds []string, sheetName string, s []string) {
	checkAndCreateList(ssIds, sheetName)
	clearData(ssIds, sheetName)
	writeStrList(ssIds, sheetName, s)
}

func clearData(ssIds []string, sheetNum string) {
	for _, ssId := range ssIds {
		_, err := srv.Spreadsheets.Values.Clear(ssId, sheetNum+"A1:N10000", &sheets.ClearValuesRequest{}).Do()
		if err != nil {
			logrus.Errorf("Unable to clearData data table: %v", err)
		} else {
			logrus.Infof("Clean sheetName %s,ssId %s", sheetNum, ssId)
		}
	}
}

func writeStrList(ssIds []string, sheetNum string, s []string) {
	var values [][]interface{}

	if len(s) > 0 {
		header := strings.Split(s[0], service.SheetSplitter)
		values = append(values, utils.ConvToISlice(header))
	} else {
		logrus.Error("The values list for writing to the sheet is empty.")
	}

	for _, line := range s[1:] {
		datum := strings.Split(line, service.SheetSplitter) //
		values = append(values, utils.ConvToISlice(datum))
	}

	valueRange := sheets.ValueRange{
		Values: values,
	}

	for _, ssId := range ssIds {
		_, err := srv.Spreadsheets.Values.
			Update(ssId, sheetNum+"A1", &valueRange).
			ValueInputOption("USER_ENTERED").
			Do()
		if err != nil {
			logrus.Errorf("Unable to write data to sheet: %v", err)
		}
	}
}

func checkAndCreateList(ssIds []string, sheetName string) {
	_sheetName := strings.TrimSuffix(sheetName, "!")
	for _, ssId := range ssIds {
		if !sheetExists(ssId, _sheetName) {
			createSheet(ssId, _sheetName)
		}
	}
}

func sheetExists(ssId string, sheetName string) bool {
	spreadsheet, err := srv.Spreadsheets.Get(ssId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve spreadsheet: %v", err)
	}

	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == sheetName {
			return true
		}
	}
	return false
}

func createSheet(ssId string, sheetName string) {
	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title: sheetName,
					},
				},
			},
		},
	}

	_, err := srv.Spreadsheets.BatchUpdate(ssId, batchUpdateRequest).Do()
	if err != nil {
		logrus.Errorf("unable to create sheet: %v", err)
	}
}
