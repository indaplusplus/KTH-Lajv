package main

//loosely based on the Jungian idea of the Self

import (
	"net/http"
)

func main() {
	http.HandleFunc("/comment/post", func(w http.ResponseWriter, r *http.Request) {
		post(w, r)
	})
	http.HandleFunc("/comment/like", func(w http.ResponseWriter, r *http.Request) {
		like(w, r)
	})
	http.HandleFunc("/comment/delete", func(w http.ResponseWriter, r *http.Request) {
		delete(w, r)
	})
	http.HandleFunc("/comment/get", func(w http.ResponseWriter, r *http.Request) {
		get(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
