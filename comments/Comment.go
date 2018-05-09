package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var database string = "http://127.0.0.1:219"

func post(w http.ResponseWriter, r *http.Request, data jsonData) {
	data.Text = filter(data.Text)
	jsonVal, err := json.Marshal(data)

	if err == nil {
		http.Post(database+"/comment", "application/json", bytes.NewBuffer(jsonVal))
	}
}

func like(w http.ResponseWriter, r *http.Request, data jsonData) {
	jsonVal, err := json.Marshal(data)

	if err == nil {
		http.Post(database+"/upvote-comment", "application/json", bytes.NewBuffer(jsonVal))
	}
}

func delete(w http.ResponseWriter, r *http.Request, data jsonData) {
	jsonVal, err := json.Marshal(data)

	if err == nil {
		http.Post(database+"/delete-comment", "application/json", bytes.NewBuffer(jsonVal))
	}
}

func get(w http.ResponseWriter, r *http.Request, data jsonData) {
	jsonVal, err1 := json.Marshal(data)

	if err1 == nil {
		resp, err2 := http.Post(database+"/get-comments", "application/json", bytes.NewBuffer(jsonVal))

		var dbData jsonData
		json.NewDecoder(resp.Body).Decode(&dbData) //decode database response

		json.NewEncoder(w).Encode(dbData) //now write out the same shit
	}
}
