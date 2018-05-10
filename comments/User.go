package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//User class, handles if they are authorized and contains info for comments

type user struct {
	Token   string `json:"token,omitempty"`
	User    string `json:"user,omitempty"`
	First   string `json:"first_name,omitempty"`
	Last    string `json:"last_name,omitempty"`
	Email   string `json:"emails,omitempty"`
	Ugkthid string `json:"ugkthid,omitempty"`
}

var loginServer string = "http://localhost:8021"

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
