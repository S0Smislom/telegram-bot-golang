package handler

import (
	"context"
	"regexp"
	"state-machine-telegram-bot/internal/adapters/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoginPhoneHandler(ctx context.Context, bot *tgbotapi.BotAPI, sm *state.StateMachine, msg *tgbotapi.Message) {
	var login string
	if msg.Contact != nil {
		login = msg.Contact.PhoneNumber
	} else {
		if match, _ := regexp.MatchString(`^(\+)[0-9]{11}$`, msg.Text); !match {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Не верно указан номер телефона"))
			return
		}
		login = msg.Text
	}

	// session.Data["login"] = login
	sm.UpdateData("login", login)
	// TODO get userr by login
	// TODO generate code and send somewhere
	// TODO save code

	// session.State = StateRegisterConfirm
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Введите код полученный из смс")
	newMsg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	bot.Send(newMsg)
	sm.SetState(state.RegisterConfirmState)
}
