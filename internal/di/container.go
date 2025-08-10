package di

import (
	botsessionstorage "state-machine-telegram-bot/internal/adapters/bot_session_storage/inmemory"
	"state-machine-telegram-bot/internal/config"
	"state-machine-telegram-bot/internal/port"
	"sync"
)

type Container struct {
	once sync.Once

	config config.Config

	botSessionStorage port.BotSessionStorage
}

func NewContainer(
	configPath string,
) (*Container, error) {
	config, err := config.LoadEnvConfig(configPath)
	if err != nil {
		return nil, err
	}
	return &Container{
		config: config,
	}, nil
}

func (s *Container) BotSessionStorage() port.BotSessionStorage {
	s.once.Do(func() {
		s.botSessionStorage = botsessionstorage.NewBotSessionStorage()
	})
	return s.botSessionStorage
}

func (s *Container) Config() config.Config {
	return s.config
}
