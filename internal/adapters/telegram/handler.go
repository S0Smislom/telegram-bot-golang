package telegram

import (
	"context"
	"log"
	"state-machine-telegram-bot/internal/adapters/telegram/handler"
	"state-machine-telegram-bot/internal/adapters/telegram/state"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *telegramBot) handleMessage(ctx context.Context, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	session := state.NewStateMachine(int(chatID), s.container.BotSessionStorage())

	if err := session.LoadSession(ctx, s.bot); err != nil {
		log.Println(err)
		return
	}
	if msg.IsCommand() {
		switch msg.Command() {
		case "start":
			handler.StartCommandHandler(ctx, s.bot, session, msg)
		case "login":
			handler.LoginCommandHandler(ctx, s.bot, session, msg)
		case "logout":
			session.UpdateData("UserID", nil)
			s.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Успешно вышел"))
			handler.StartCommandHandler(ctx, s.bot, session, msg)
		}
		return
	}
	switch session.GetState() {
	case state.BaseState:
		switch msg.Text {
		case "Календарь":
			handler.SendCalendar(ctx, s.bot, session, msg)
		default:
			handler.DefaultHandler(ctx, s.bot, msg)
		}
	case state.RegisterState:
		switch msg.Text {
		case "Назад":
			handler.StartCommandHandler(ctx, s.bot, session, msg)
		default:
			handler.LoginPhoneHandler(ctx, s.bot, session, msg)

		}
	case state.RegisterConfirmState:
		switch msg.Text {
		case "Назад":
			handler.LoginCommandHandler(ctx, s.bot, session, msg)

		default:
			handler.LoginCodeHandler(ctx, s.bot, session, msg)
		}
	default:
		handler.StartCommandHandler(ctx, s.bot, session, msg)
	}

}

func (s *telegramBot) handleCallbackQuery(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	if _, err := s.bot.Request(tgbotapi.NewCallback(callback.ID, "")); err != nil {
		log.Println("Callback error:", err)
	}
	switch {
	case strings.HasPrefix(callback.Data, "button::"):
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "BUTTON:"+callback.Data)
		s.bot.Send(msg)
	case strings.HasPrefix(callback.Data, "link::"):
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
		s.bot.Send(msg)
	}
}
