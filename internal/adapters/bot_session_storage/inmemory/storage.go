package botsessionstorage

import (
	"context"
	"errors"
	"state-machine-telegram-bot/internal/port"
)

type botSessionStorate struct {
	store map[string][]byte
}

func NewBotSessionStorage() port.BotSessionStorage {
	return &botSessionStorate{
		store: make(map[string][]byte),
	}
}

func (s *botSessionStorate) Set(ctx context.Context, key string, data []byte) error {
	s.store[key] = data
	return nil
}
func (s *botSessionStorate) Get(ctx context.Context, key string) ([]byte, error) {
	v, ok := s.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return v, nil
}
