package main

import (
	"errors"
	ws "github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type StreamID int

type Chat struct {
	channels map[StreamID]*Channel
	db       Database
}

func NewChat() Chat {
	return Chat{
		make(map[StreamID]*Channel),
		Database{},
	}
}

func authenticate(v url.Values) (string, error) {
	names, ok := v["name"]
	if !ok && len(names) < 1 {
		return "", errors.New("No username provided")
	}
	tokens, ok := v["token"]
	if !ok && len(tokens) < 1 {
		return "", errors.New("No token provided")
	}
	name := names[0]
	token := tokens[0]
	if name == "filip" && token == "token" {
		return name, nil
	}
	return "", errors.New("Incorrect auth details")
}

func getStreamId(v url.Values) (StreamID, error) {
	strIDs, ok := v["sid"]
	if !ok && len(strIDs) < 1 {
		return 0, errors.New("No Stream ID provided")
	}
	sid, err := strconv.Atoi(strIDs[0])
	if err != nil {
		return 0, errors.New("Could not parse Stream ID")
	}
	return StreamID(sid), nil
}

func (c *Chat) getChannel(sid StreamID) *Channel {
	//TODO: Fix race-condition on the map here.
	if chn, ok := c.channels[sid]; ok {
		return chn
	}
	// If two people connect simaltanious to a channel that does not exist
	// they will both create a new channel and one will be placed in limbo.
	// Possibly kept alive by the client connected to it so it could create
	// a memmory leak too.
	nchn := NewChannel()
	nchn.run()
	c.channels[sid] = &nchn
	return &nchn
}

func (c Chat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	uname, err := authenticate(q)
	if err != nil {
		log.Println("User auth failed")
		log.Println("error:", err)
		return
	}
	sid, err := getStreamId(q)
	if err != nil {
		log.Println("error:", err)
		return
	}

	chn := c.getChannel(sid)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("success: Connecting", uname, "to channel", sid)
	chn.subscribe(uname, conn)
}
