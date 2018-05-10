package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func SendStreamRequest(url string, jsonData *StreamData) []byte {
	jsonString, err := json.Marshal(jsonData)
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
