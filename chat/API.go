package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type API struct{}

func getUser(name, token string) (string, error) {
	// PORT: 8021
	// http://localhost:8021/getUser?token={token}

	q, err := url.Parse("http://localhost:8021/getUser")
	if err != nil {
		return false
	}
	q := u.Query()
	q.Set("token", token)

	log.Println("Sending request to", q.String())
	res, err := http.Get(q.String())
	if err != nil {
		return "", err
	}

	user, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

type queryObject struct {
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
}

func queryDB(data queryObject) (string, error) {
	res, err := http.Post("http://localhost:55994", "application/json", json.Marshal(data))
	if err != nil {
		return "", errors.New("Could not querry database")
	}
}

func (d API) getStreamChats(sid StreamID) ([]Message, error) {
	return []Message{}, errors.New("NOT IMPLEMENTED")
}

func (a API) storeMessage(m Message) error {

}

func streamIsLive(sid StreamID) error {

}
