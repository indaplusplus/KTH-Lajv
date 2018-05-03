package main

import (
	"encoding/json"
	ws "github.com/gorilla/websocket"
	"log"
	"strings"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024

	// The number of message to buffer in a channel
	messageBufferSize = 64
)

type Client struct {
	// Reference to channel
	ch *Channel

	// WebSocket
	conn *ws.Conn

	// A screen name for the pressenting the sender.
	Name string

	// A buffered cannel for messages to be sent to the user.
	send chan Message
}

func NewClient(uname string, conn *ws.Conn, ch *Channel) Client {
	return Client{
		ch:   ch,
		conn: conn,
		Name: uname,
		send: make(chan Message, messageBufferSize),
	}
}

func (c *Client) run() {
	c.ch.register <- c
	go c.readRoutine()
	go c.writeRoutine()
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readRoutine() {
	defer func() {
		c.ch.unregister <- c
		c.conn.Close()
	}()

	// Limit the size of the messages we read
	c.conn.SetReadLimit(maxMessageSize) // TODO: Move to creation of conn
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			} // Else the Close was expected aka no error.
			return
		}

		message := NewMessage(c.Name, strings.Trim(string(msg), " \n\r"))
		c.ch.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writeRoutine() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(ws.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(ws.TextMessage)
			if err != nil {
				return
			}
			msg, err := json.Marshal(message)

			if err != nil {
				log.Println("error:", err)
			}

			w.Write(msg)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(ws.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
