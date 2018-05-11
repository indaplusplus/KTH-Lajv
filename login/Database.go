package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
)

var DB string = "http://localhost:55994"
// var db map[string]string = map[string]string{"4T0k3n": "filip", "token": "usr"}

func loginToken(token string, user string) {
	//db[token] = user
	data, _ := json.Marshal(jsonData{Command: "login", Token: token, User: user})
	_, err := http.Post(DB, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func logoutToken(token string) {
	// delete(db, token)
	data, _ := json.Marshal(jsonData{Command: "logout", Token: token})
	_, err := http.Post(DB, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func containsToken(token string) bool {
	_, has := getLoggedInUser(token)
	return has
}

func getLoggedInUser(token string) (string, bool) {
	// user, has := db[token]
	data, _ := json.Marshal(jsonData{Command: "logout", Token: token})
	resp, err := http.Post(DB, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	var respData jsonData
	json.NewDecoder(resp.Body).Decode(&respData)
	return respData.User, respData.User != ""
}
