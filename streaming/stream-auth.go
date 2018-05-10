package main

import (
	"encoding/json"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Key struct {
        Id int `json:"id"`
        Token string `json:"Token"`
}

func AuthStreamHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	name := v.Get("name")

	key_json_str, err_b64 := base64.StdEncoding.DecodeString(name)
	if err_b64 != nil {
		panic(err_b64)
	}

	var key Key
	err := json.Unmarshal(key_json_str, &key)
	if err != nil {
		panic(err)
	}

	// TODO: check if token is valid
	//return_code := http.StatusUnauthorized
	//return_msg := "Invalid stream key"

	//stop-stream works more like a update than a stop...
	update_json := StreamData{
		Command: "stop-stream",
		Vod:     "",
		Hls:     "http://live.edstrom.me:6060/hls/" + strconv.Itoa(key.Id) + ".m3u8",
		// Stream is only needed for the streamer, this should be fine.
		Stream: "",
		Id:     key.Id,
	}
	_ = SendStreamRequest("http://localhost:55994/", &update_json)

	return_code := http.StatusOK
	return_msg := "OK"

	w.WriteHeader(return_code)
	fmt.Fprintf(w, return_msg)
}

var HlsLocation string = "/tmp/hls"

func DoneStreamHandler(w http.ResponseWriter, r *http.Request) {
	// Vods are a work in progress, encoding from FLV to mp4 needs to be done.
	// Reset Hls, it's not needed anymore..
	v := r.URL.Query()
	name := v.Get("name")

	key_json_str, err := base64.StdEncoding.DecodeString(name)
	if err != nil {
		panic(err)
	}

	var key Key
	err = json.Unmarshal(key_json_str, &key)
	if err != nil {
		panic(err)
	}

	// Note: It's probably redundant checking for token here.
	stop_json := StreamData{
		Command: "stop-stream",
		Vod:     strconv.Itoa(key.Id) + ".mp4",
		Hls:     "",
		Stream:  "",
		Id:      key.Id,
	}

	var data StreamData
	res := SendStreamRequest("http://localhost:55994/", &stop_json)
	err = json.Unmarshal(res, &data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

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
