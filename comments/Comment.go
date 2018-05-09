package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var database string = "http://127.0.0.1:55994"

func post(w http.ResponseWriter, r *http.Request, data jsonData) {
	data.Command = "comment"
	data.Text = filter(data.Text)
	jsonVal, err := json.Marshal(data)

	if err != nil {
		return
	}

	http.Post(database, "application/json", bytes.NewBuffer(jsonVal))
}

func like(w http.ResponseWriter, r *http.Request, data jsonData) {
	data.Command = "upvote-comment"
	jsonVal, err := json.Marshal(data)

	if err != nil {
		return
	}

	http.Post(database, "application/json", bytes.NewBuffer(jsonVal))
}

func delete(w http.ResponseWriter, r *http.Request, data jsonData) {
	data.Command = "delete-comment"
	jsonVal, err := json.Marshal(data)

	if err != nil {
		return
	}

	http.Post(database, "application/json", bytes.NewBuffer(jsonVal))
}

func get(w http.ResponseWriter, r *http.Request, data jsonData) {
	data.Command = "get-comments"
	jsonVal, err1 := json.Marshal(data)

	if err1 != nil {
		return
	}

	resp, _ := http.Post(database, "application/json", bytes.NewBuffer(jsonVal))

	var dbData jsonData
	json.NewDecoder(resp.Body).Decode(&dbData)

	json.NewEncoder(w).Encode(dbData.Comments)
}
