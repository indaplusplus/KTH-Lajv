package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var database string = "http://127.0.0.1:291"

func post(w http.ResponseWriter, r *http.Request, data jsonData) {
	jsonVal, err := json.Marshal(data)

	if err == nil {

	}

	http.Post(database+"/comment", "application/json", bytes.NewBuffer(jsonVal))
}

func like(w http.ResponseWriter, r *http.Request, data jsonData) {
}

func delete(w http.ResponseWriter, r *http.Request, data jsonData) {
}

func get(w http.ResponseWriter, r *http.Request, data jsonData) {
	jsonVal, err1 := json.Marshal(data)

	if err1 == nil {

	}

	resp, err2 := http.Post(database+"/get-comments", "application/json", bytes.NewBuffer(jsonVal))

	var dbData jsonData
	json.NewDecoder(r.Body).Decode(&dbData)

	//pass this back to the client
}
