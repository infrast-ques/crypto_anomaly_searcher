package telegram

import (
	"crypto_anomaly_searcher/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var TgBot = initTgBot()

type TgClient struct {
	Bot *tgbotapi.BotAPI
}

func initTgBot() TgClient {
	bot, err := tgbotapi.NewBotAPI(common.ConfigData.Telegram.ApiKey)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	return TgClient{Bot: bot}
}

// func MessageHandler(tg TgClient) {
// 	updateConfig := tgbotapi.NewUpdate(0)
// 	updateConfig.Timeout = 5
// 	updates := tg.Bot.GetUpdatesChan(updateConfig)
// 	for upd := range updates {
// 		if upd.Message == nil {
// 			continue
// 		}
// 		service.Logger.Infof("Received a message from userID %d", upd.Message.From.ID)
// 		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
// 		if _, err := tg.Bot.Send(msg); err != nil {
// 			service.Logger.Errorf("Can`t send reply to userId = %x", upd.Message.MessageID)
// 		}
// 	}
// }

func (tg TgClient) SendLog(s string) (tgbotapi.Message, error) {
	return tg.SendMsg(s, common.ConfigData.Telegram.TestUserId)
}

func (tg TgClient) SendMsg(data string, userId int64) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(userId, data)
	return tg.Bot.Send(msg)
}
