package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var LOGIN_API_URL string = "https://login2.datasektionen.se"
var LOGIN_COMPLETE_URL string = "http://" + GetOutboundIP() + ":8021/loginComplete"
var LOGIN_API_KEY string = os.Getenv("LOGIN_API_KEY")

func login(w http.ResponseWriter, r *http.Request) {
	url := LOGIN_API_URL + "/login?callback=" + LOGIN_COMPLETE_URL + "?token="
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func loginComplete(w http.ResponseWriter, r *http.Request) {
	var token string = getTokenFromURL(r)
	url := LOGIN_API_URL + "/verify/" + token + ".json?api_key=" + LOGIN_API_KEY
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	var tmp jsonData
	json.NewDecoder(resp.Body).Decode(&tmp)
	if tmp.User == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	loginToken(token, tmp.User)
	json.NewEncoder(w).Encode(jsonData{Token: token})
}

func logout(w http.ResponseWriter, r *http.Request) {
	token := getTokenFromJson(r)
	logoutToken(token)
	w.WriteHeader(http.StatusOK)
}

func isLoggedin(w http.ResponseWriter, r *http.Request) {
	token := getTokenFromJson(r)
	if containsToken(token) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	token := getTokenFromJson(r)
	user, has := getLoggedInUser(token)
	if has {
		json.NewEncoder(w).Encode(jsonData{User: user})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func addDummyData(w http.ResponseWriter, r *http.Request) {
	loginToken("4T0k3n", "filip")
	loginToken("token", "usr")
}

func register(w http.ResponseWriter, r *http.Request) {
	//TODO if needed
}

func getTokenFromJson(r *http.Request) string {
	var data jsonData
	json.NewDecoder(r.Body).Decode(&data)
	return data.Token
}

func getTokenFromURL(r *http.Request) string {
	token, has := r.URL.Query()["token"]
	if !has || len(token) < 1 {
		return ""
	}
	return token[0]
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
