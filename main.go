package main

import (
	"crypto_anomaly_searcher/common"
	"crypto_anomaly_searcher/service/data_collector"
	"crypto_anomaly_searcher/service/scheduler"
	"crypto_anomaly_searcher/service/sheets"
)

func main() {
	scheduler.ScheduleTask(func() {
		data := data_collector.AggregateData()
		sheets.FillSheet(common.ConfigData.Sheet.SsIds, common.SheetListRawData, data.RawData.ToStringList())
		sheets.FillSheet(common.ConfigData.Sheet.SsIds, common.SheetListComputed1, data.Computed1Data)
	})

	// bot := telegram.InitTgBot()
	// telegram.SendToUsers(utils.StrListToStr(data), bot)
}
