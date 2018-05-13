package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
}

type messageData struct {
	User        string `json:"user"`
	Time        string `json:"time"`
	Text        string `json:"text"`
	Upvotes     int    `json:"upvotes"`
	ReplyToUser string `json:"replyToUser"`
	ReplyToTime string `json:"replyToTime"`
}

type StreamData struct {
	Command  string `json:"command"`
	Course   string `json:"course"`
	Room     string `json:"room"`
	Lecturer string `json:"lecturer"`
	Streamer string `json:"streamer"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Vod      string `json:"vod"`
	Stream   string `json:"stream"`
	Hls      string `json:"hls"`
	Id       int    `json:"id"`
	Ids      []int  `json:"ids"`
	Token    string `json:"token"`
	Loggedin bool   `json:"loggedin"`
}

func QueryLoggedIn(auth_token string) *StreamData {
	json_auth_check := StreamData{
		Command: "loggedin",
		Token:   auth_token,
	}

	var data StreamData
	res := SendStreamRequest("http://localhost:219/", &json_auth_check)
	err := json.Unmarshal(res, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func unpackJson(r *http.Request) *StreamData {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var data StreamData
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func SendJSONRequest(url string, ptr_json *jsonData) []byte {
	jsonString, err := json.Marshal(ptr_json)
	if err != nil {
		panic(err)
	}
	data := []byte(jsonString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err_read := ioutil.ReadAll(resp.Body)
	if err_read != nil {
		panic(err_read)
	}
	return body
}

func SendStreamRequest(url string, ptr_json *StreamData) []byte {
	jsonString, err := json.Marshal(ptr_json)
	if err != nil {
		panic(err)
	}
	data := []byte(jsonString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err_read := ioutil.ReadAll(resp.Body)
	if err_read != nil {
		panic(err_read)
	}
	return body
}
