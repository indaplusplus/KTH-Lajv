package main

//loosely based on the Jungian idea of the Self

import (
	"encoding/json"
	"net/http"
)

type jsonData struct {
	Command     string        `json:"command"`
	Course      string        `json:"course"`
	Room        string        `json:"room"`
	Lecturer    string        `json:"lecturer"`
	Streamer    string        `json:"streamer"`
	Name        string        `json:"name"`
	Date        string        `json:"date"`
	Vod         string        `json:"vod"`
	Stream      string        `json:"stream"`
	Hls         string        `json:"hls"`
	Id          int           `json:"id"`
	Ids         []int         `json:"ids"`
	User        string        `json:"user"`
	Time        string        `json:"time"`
	Text        string        `json:"text"`
	ReplyToUser string        `json:"replyToUser"`
	ReplyToTime string        `json:"replyToTime"`
	Chat        []messageData `json:"chat"`
	Comments    []messageData `json:"comments"`
	Token       string        `json:"token"`
	Loggedin    bool          `json:"loggedin"`
}

type messageData struct {
	User        string `json:"user"`
	Time        string `json:"time"`
	Text        string `json:"text"`
	Upvotes     int    `json:"upvotes"`
	ReplyToUser string `json:"replyToUser"`
	ReplyToTime string `json:"replyToTime"`
}

func main() {
	http.HandleFunc("/comment/post", func(w http.ResponseWriter, r *http.Request) {
		post(w, r, getData(r))
	})
	http.HandleFunc("/comment/like", func(w http.ResponseWriter, r *http.Request) {
		like(w, r, getData(r))
	})
	http.HandleFunc("/comment/delete", func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, getData(r))
	})
	http.HandleFunc("/comment/get", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, getData(r))
	})

	http.ListenAndServe(":8080", nil)
}

func getData(r *http.Request) jsonData {
	var data jsonData
	json.NewDecoder(r.Body).Decode(&data)
	return data
}
