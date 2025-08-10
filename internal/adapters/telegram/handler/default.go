package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DefaultHandler(ctx context.Context, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	newMessage := tgbotapi.NewMessage(msg.Chat.ID, "Не понял")
	newMessage.ReplyToMessageID = msg.MessageID
	bot.Send(newMessage)
}
