package main

import "crypto_anomaly_searcher/my_service"

func main() {
	data := my_service.GetCryptoData()
	my_service.WriteData(my_service.ToSheetData(data))
}
