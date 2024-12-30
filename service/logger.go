package service

import (
	"io"

	"crypto_anomaly_searcher/service/telegram"
	"github.com/sirupsen/logrus"
)

func setUpLogger(logger *logrus.Logger, wr io.Writer) *logrus.Logger {
	logger.SetOutput(wr)
	return logger
}

// Logger todo переделать как то по нормальному
var Logger = setUpLogger(logrus.New(), huinya)

var huinya = TgClientWrapper{
	hui: telegram.TgBot,
}

type TgClientWrapper struct {
	hui telegram.TgClient
}

func (tg TgClientWrapper) Write(p []byte) (n int, err error) {
	_, err = tg.hui.SendLog(string(p))
	return 0, err
}
