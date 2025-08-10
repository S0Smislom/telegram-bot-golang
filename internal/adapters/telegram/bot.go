package telegram

import (
	"context"
	"log"
	"os"
	"os/signal"
	"state-machine-telegram-bot/internal/di"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramBot struct {
	container *di.Container
	bot       *tgbotapi.BotAPI
}

func NewTelegramBotV2(configPath string) (*telegramBot, error) {
	container, err := di.NewContainer(configPath)
	if err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(container.Config().TelegramBotToken)
	if err != nil {
		return nil, err
	}

	return &telegramBot{
		container: container,
		bot:       bot,
	}, nil
}

func (s *telegramBot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.bot.GetUpdatesChan(u)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("FINISHED")
			return nil
		case update := <-updates:
			if update.Message != nil {
				s.handleMessage(ctx, update.Message)
			}
			if update.CallbackQuery != nil {
				s.handleCallbackQuery(ctx, update.CallbackQuery)
			}
		}
	}
}
