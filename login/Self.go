package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type jsonData struct {
	Token   string `json:"token,omitempty"`
	User    string `json:"user,omitempty"`
	First   string `json:"first_name,omitempty"`
	Last    string `json:"last_name,omitempty"`
	Email   string `json:"emails,omitempty"`
	Ugkthid string `json:"ugkthid,omitempty"`
	Command string `json:"command,omitempty"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("GET")
	router.HandleFunc("/loginComplete", loginComplete).Methods("GET")
	router.HandleFunc("/logout", logout).Methods("POST")
	router.HandleFunc("/isLoggedin", isLoggedin).Methods("POST")
	router.HandleFunc("/getUser", getUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8021", router))
}
