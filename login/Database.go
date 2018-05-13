package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var DB string = "http://localhost:55994/"

func loginToken(token string, user string) {
	fmt.Printf("Adding (%s, %s) to db\n", token, user)
	data, _ := json.Marshal(jsonData{Command: "login", Token: token, User: user})
	_, err := http.Post(DB, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func logoutToken(token string) {
	fmt.Println("removing token", token, "from db")
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
	fmt.Println("retrieving token", token, "from db")
	data, _ := json.Marshal(jsonData{Command: "loggedin", Token: token})
	resp, err := http.Post(DB, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	var respData jsonData
	json.NewDecoder(resp.Body).Decode(&respData)
	return respData.User, respData.User != ""
}
