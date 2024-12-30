package scheduler

import (
	"time"

	"crypto_anomaly_searcher/common"
	"crypto_anomaly_searcher/service"
)

func ScheduleTask(task func()) {
	service.Logger.Info("Execute task")
	config := common.ConfigData.Sheet
	task()
	ticker := time.NewTicker(time.Duration(config.UpdateTime) * time.Minute)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				service.Logger.Info("Execute task")
				task()
			case <-done:
				return
			}
		}
	}()

	time.Sleep(240 * time.Hour)
	done <- true
	service.Logger.Info("Stop scheduler")
}
