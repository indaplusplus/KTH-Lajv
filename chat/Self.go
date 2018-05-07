package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", NewChat())
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}