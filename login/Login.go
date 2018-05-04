package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

var LOGIN_API_URL string = "https://login2.datasektionen.se"
var LOGIN_COMPLETE_URL string = "http://localhost:8050/login/loginComplete"

type Response struct {
	Resp  string `json:"isLoggedin,omitempty"`
	Token string `json:"token,omitempty"`
}

func login(w http.ResponseWriter, r *http.Request) {
	url := LOGIN_API_URL + "/login?callback=" + LOGIN_COMPLETE_URL + "?token="
  http.Redirect(w, r, url, http.StatusSeeOther)
}

func loginComplete(w http.ResponseWriter, r *http.Request) {
	var token string = getToken(r);
	//TODO: Verify token
	loginToken(token)
	json.NewEncoder(w).Encode(Response{Token: token})
}

func logout(w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	logoutToken(token)
}

func isLoggedin(w http.ResponseWriter, r *http.Request) {
	var ret Response
	token := getToken(r)
	if containsToken(token) {
		ret.Resp = "true"
	} else {
		ret.Resp = "false"
	}
	json.NewEncoder(w).Encode(ret)
}

func register(w http.ResponseWriter, r *http.Request) {
	//TODO if needed
}

func getToken(r *http.Request) string {
	return string(regexp.MustCompile(`token=\w{22}`).Find([]byte(fmt.Sprint(r.URL))))[6:]
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login/login", login).Methods("GET")
	router.HandleFunc("/login/loginComplete", loginComplete).Methods("GET")
	router.HandleFunc("/login/logout", logout).Methods("GET")
	router.HandleFunc("/login/isLoggedin", isLoggedin).Methods("GET")
	log.Fatal(http.ListenAndServe(":8050", router))
}
