package state

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"telegram-bot/internal/port"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type State string

const (
	BaseState            State = "base"
	RegisterState        State = "register"
	RegisterConfirmState State = "register-confirm"
)

func ParseState(s string) (State, error) {
	switch State(s) {
	case BaseState, RegisterState, RegisterConfirmState:
		return State(s), nil
	default:
		return "", errors.New("wrong state")
	}
}

// type State interface {
// 	Execute(ctx context.Context, msg *tgbotapi.Message) error
// 	GetName() State
// }

type StateMachine struct {
	chatID int
	data   map[string]any

	state   State
	storage port.BotSessionStorage
}

func NewStateMachine(chatID int, sessionStorate port.BotSessionStorage) *StateMachine {
	return &StateMachine{
		chatID:  chatID,
		storage: sessionStorate,
	}
}

func (s *StateMachine) LoadSession(ctx context.Context, bot *tgbotapi.BotAPI) error {
	sessionStr, err := s.storage.Get(context.Background(), strconv.Itoa(s.chatID))
	if err != nil {
		s.SetState(BaseState)
		return nil
	}

	var sessionMap map[string]any
	if err := json.Unmarshal(sessionStr, &sessionMap); err != nil {
		return err
	}

	state, err := ParseState(sessionMap["state"].(string))
	if err != nil {
		state = BaseState
	}
	s.SetState(state)

	if sessionMap != nil {
		if v := sessionMap["data"]; v != nil {
			s.SetData(v.(map[string]any))
		}
	}
	return nil
}

func (s *StateMachine) GetState() State {
	return s.state
}

func (s *StateMachine) SetState(state State) {
	s.state = state
	s.save()
}

func (s *StateMachine) UpdateData(key string, value any) {
	if s.data == nil {
		s.data = map[string]any{}
	}
	s.data[key] = value
	s.save()
}

func (s *StateMachine) SetData(data map[string]any) {
	s.data = data
	s.save()
}

func (s *StateMachine) GetData(key string) (any, bool) {
	v, ok := s.data[key]
	return v, ok
}

// func (s *StateMachine) Execute(ctx context.Context, msg *tgbotapi.Message) {
// 	s.state.Execute(ctx, msg)
// }

func (s *StateMachine) save() {
	payload := map[string]any{}

	payload["state"] = string(s.state)
	payload["data"] = s.data
	js, _ := json.Marshal(payload)

	s.storage.Set(context.Background(), strconv.Itoa(s.chatID), js)
}
