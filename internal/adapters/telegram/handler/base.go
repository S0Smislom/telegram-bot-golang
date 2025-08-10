package handler

import (
	"context"
	"fmt"
	"telegram-bot/internal/adapters/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommandHandler(ctx context.Context, bot *tgbotapi.BotAPI, sm *state.StateMachine, msg *tgbotapi.Message) {
	var textToSend string
	if v, _ := sm.GetData("UserID"); v != nil {
		textToSend = fmt.Sprintf("Я тебя знаю %v", v)
	} else {
		textToSend = "Привет, как твои дела, бро"
	}
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, textToSend)
	// newMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	newMsg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Календарь"),
		),
	)
	bot.Send(newMsg)

	sm.SetState(state.BaseState)
}

func LoginCommandHandler(ctx context.Context, bot *tgbotapi.BotAPI, sm *state.StateMachine, msg *tgbotapi.Message) {
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Введите номер телефона")
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Поделиться номером телефона"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	newMsg.ReplyMarkup = buttons
	buttons.ResizeKeyboard = true
	bot.Send(newMsg)
	sm.SetState(state.RegisterState)
}

func SendCalendar(ctx context.Context, bot *tgbotapi.BotAPI, sm *state.StateMachine, msg *tgbotapi.Message) {
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Календарь")
	newMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("A", "button::a"),
			tgbotapi.NewInlineKeyboardButtonData("B", "button::b"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Link", "link::a"),
		),
	)
	bot.Send(newMsg)
}
