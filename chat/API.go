package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type user struct {
	Token   string `json:"token,omitempty"`
	User    string `json:"user,omitempty"`
	First   string `json:"first_name,omitempty"`
	Last    string `json:"last_name,omitempty"`
	Email   string `json:"emails,omitempty"`
	Ugkthid string `json:"ugkthid,omitempty"`
}

var loginServer = "http://localhost:8021"

func loggedIn(token string) bool {
	u := user{Token: token}
	jsonVal, err1 := json.Marshal(u)

	if err1 != nil {
		return false
	}

	resp, err2 := http.Post(loginServer+"/isLoggedin", "application/json", bytes.NewBuffer(jsonVal))

	if err2 != nil {
		return false
	}

	return resp.StatusCode == 200
}

func getUsername(token string) string {
	u := user{Token: token}
	jsonVal, err1 := json.Marshal(u)

	if err1 != nil {
		return ""
	}

	resp, err2 := http.Post(loginServer+"/getUser", "application/json", bytes.NewBuffer(jsonVal))

	if err2 != nil || resp.StatusCode == 401 {
		return ""
	}

	var res user
	json.NewDecoder(resp.Body).Decode(&res)
	return res.User
}

var database = "http://localhost:55994"

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

func storeMessage(m Message, sid StreamID) error {
	var data jsonData
	data.Command = "chat"
	data.Id = int(sid)
	msg, err := json.Marshal(m)
	data.Text = string(msg)
	jsonVal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	http.Post(database, "application/json", bytes.NewBuffer(jsonVal))
	return nil
}
