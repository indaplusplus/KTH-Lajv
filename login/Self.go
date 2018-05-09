package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("GET")
	router.HandleFunc("/loginComplete", loginComplete).Methods("GET")
	router.HandleFunc("/logout", logout).Methods("POST")
	router.HandleFunc("/isLoggedin", isLoggedin).Methods("POST")
	router.HandleFunc("/getUser", getUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8021", router))
}
