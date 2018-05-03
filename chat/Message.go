package main

import (
	"time"
)

func getNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type Message struct {
	// Who sent it
	From string `json:from`

	// The content of the message
	Text string `json:text`

	// Time it was sent (recived at the server)
	Time int64 `json:time`
}

func NewMessage(user string, message string) Message {
	return Message{
		user,
		message,
		getNow(),
	}
}
