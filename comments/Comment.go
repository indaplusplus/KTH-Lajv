package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type comment struct {
	id       int       `json:"id"`
	user     string    `json:"user"`
	time     time.Time `json:"time"`
	text     string    `json:"text"`
	upvotes  int       `json:"upvotes"`
	response int       `json:"response"`
}

func post(w http.ResponseWriter, r *http.Request) { //takes user id, token, text, response
	m := getMap(r)

	user := user{id: m["user"].(string), token: m["token"].(string)}
	text := m["text"].(string)
	response := m["text"].(int)
}

func like(w http.ResponseWriter, r *http.Request) { //takes user id, token, post id
	m := getMap(r)

	user := user{id: m["user"].(string), token: m["token"].(string)}
	postid := m["postid"].(int)
}

func delete(w http.ResponseWriter, r *http.Request) { //takes user id, token, post id
	m := getMap(r)

	user := user{id: m["user"].(string), token: m["token"].(string)}
	postid := m["postid"].(int)
}

func get(w http.ResponseWriter, r *http.Request) { //takes user id, token, video
	m := getMap(r)

	user := user{id: m["user"].(string), token: m["token"].(string)}
	video := m["video"].(string)

	//fetch comments associated with video

	//send comments associated with video
}

func getMap(r *http.Request) map[string]interface{} {
	var msg interface{}
	json.NewDecoder(r.Body).Decode(&msg)
	return msg.(map[string]interface{})
}
