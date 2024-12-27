package main

import (
	"crypto_anomaly_searcher/service"
)

func main() {
	data := service.AggregateData()
	service.ClearSheet()
	service.WriteData(data)

	// bot := telegram.InitTgBot()
	// telegram.SendToUsers(utils.StrListToStr(data), bot)
}
