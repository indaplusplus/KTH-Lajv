package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RequestStreamPermission(name string) bool {
	if name == "lajv" {
		return true
	}
	return false
}

func AuthStreamHandler(w http.ResponseWriter, r *http.Request) {
	//name is our stream key sent from nginx
	v := r.URL.Query()
	name := v.Get("name")

	return_code := http.StatusUnauthorized
	return_msg := "Invalid stream key"

	if name != "" {
		if RequestStreamPermission(name) {
			return_code = http.StatusOK
			return_msg = "OK"
		}
	}

	w.WriteHeader(return_code)
	fmt.Fprintf(w, return_msg)
}

func DoneStreamHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", AuthStreamHandler).Methods("GET")
	r.HandleFunc("/done", DoneStreamHandler).Methods("GET")
	http.Handle("/", r)
	server := &http.Server{
		Handler: r,
		Addr:    "localhost:1337",
	}
	log.Fatal(server.ListenAndServe())
}
