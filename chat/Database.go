package main

import (
	"errors"
)

type Database struct{}

func (d Database) userOk(name, token string) (bool, error) {
	return true, errors.New("NOT IMPLEMENTED")
}

func (d Database) getStreamChats(sid StreamID) ([]Message, error) {
	return []Message{}, errors.New("NOT IMPLEMENTED")
}
