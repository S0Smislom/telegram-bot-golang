package main

import "context"

type StateType string
type Event string

type Action interface {
	Execute(ctx context.Context) StateType
}

type State struct {
	Action Action
	Events []Event
}

type StateMachine struct {
	currentState State
	States       map[StateType]State
}

/*
	Что надо сделать?
	Нужен объек,который будет знать о всех состояниях, переходах между состояниями и событиями

*/
