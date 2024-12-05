package main

import "crypto_anomaly_searcher/my_service"

func main() {
	data := my_service.AggregateData()
	my_service.ClearSheet()
	my_service.WriteData(my_service.ToSheetData(data))
}
