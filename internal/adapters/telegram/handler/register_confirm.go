package handler

import (
	"context"
	"fmt"
	"regexp"
	"state-machine-telegram-bot/internal/adapters/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoginCodeHandler(ctx context.Context, bot *tgbotapi.BotAPI, sm *state.StateMachine, msg *tgbotapi.Message) {
	if match, _ := regexp.MatchString(`^[0-9]{4}$`, msg.Text); !match {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Не верно указан код"))
		return
	}

	// TODO logic to authenticate user
	login, _ := sm.GetData("login")
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Успешный успех: %s %s", login, msg.Text))
	newMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(newMsg)

	sm.UpdateData("UserID", 1)

	StartCommandHandler(ctx, bot, sm, msg)

}
