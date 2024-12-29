package scheduler

import (
	"time"

	"github.com/sirupsen/logrus"
)

func ScheduleTask(task func()) {
	logrus.Info("Execute task")
	task()
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				logrus.Info("Execute task")
				task()
			case <-done:
				return
			}
		}
	}()

	time.Sleep(30 * time.Minute)
	done <- true
	logrus.Info("Stop scheduler")
}
