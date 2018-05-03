package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", NewChat())
	fmt.Println("Starting server on :9876")
	http.ListenAndServe(":9876", nil)
}
