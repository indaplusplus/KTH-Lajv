package main

import (
	"net/http"
	"time"
)

type comment struct {
	id       int        `json:"id"`
	user     string     `json:"user"`
	time     time.Stamp `json:"time"`
	text     string     `json:"text"`
	upvotes  int        `json:"upvotes"`
	response int        `json:"response"`
}

func post(w http.ResponseWriter, r *http.Request) { //takes user id, token, text, response

}

func like(w http.ResponseWriter, r *http.Request) { //takes user id, token, post id

}

func delete(w http.ResponseWriter, r *http.Request) { //takes user id, token, post id

}

func get(w http.ResponseWriter, r *http.Request) { //takes user id, token, video
	//fetch comments associated with video
}
