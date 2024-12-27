package telegram

import (
	"crypto_anomaly_searcher/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func InitTgBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(service.ConfigData.Telegram.ApiKey)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	return bot
}

func MessageHandler(bot *tgbotapi.BotAPI) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 5
	updates := bot.GetUpdatesChan(updateConfig)
	for upd := range updates {
		if upd.Message == nil {
			continue
		}
		logrus.Infof("Received a message from userID %d", upd.Message.From.ID)
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
		if _, err := bot.Send(msg); err != nil {
			logrus.Errorf("Can`t send reply to userId = %x", upd.Message.MessageID)
		}
	}
}

func SendToUsers(data string, bot *tgbotapi.BotAPI) {

	msg := tgbotapi.NewMessage(service.ConfigData.Telegram.TestUserId, data)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Алиса", "alice"),
			tgbotapi.NewInlineKeyboardButtonData("Боб", "bob"),
		))
	bot.Send(msg)

}
