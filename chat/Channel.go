package main

import (
	ws "github.com/gorilla/websocket"
)

type Channel struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func NewChannel() Channel {
	return Channel{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (c *Channel) run() {
    go func() {
        for {
            select {
            case client := <-c.register:
                c.clients[client] = true
            case client := <-c.unregister:
                if _, ok := c.clients[client]; ok {
                    delete(c.clients, client)
                    close(client.send)
                }
            case message := <-c.broadcast:
                for client := range c.clients {
                    select {
                    case client.send <- message:
                    default:
                        close(client.send)
                        delete(c.clients, client)
                    }
                }
            }
        }
    }()
}

func (c *Channel) subscribe(uname string, conn *ws.Conn) {
	// Create the client
	client := NewClient(uname, conn, c)
	client.run()
}
