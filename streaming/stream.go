package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"encoding/base64"
	"log"
	"net/http"
)

type CreateStreamStruct struct {
	Course   string `json:"course"`
	Room     string `json:"room"`
	Streamer string `json:"streamer"`
	Lecturer string `json:"lecturer"`
	Name     string `json:"name"`
}

type CreateStreamResult struct {
	Key     string    `json:"key"`
	Stream string `json:"stream"`
}

type Key struct {
	Id int `json:"id"`
	Token string `json:"Token"`
}

// The key is pretty much a base64:ed JSON object containing both stream id and the users token.
func MakeKey(streamId int, token string)(string) {
	key := Key{
		Id: streamId,
		Token: token,
	}
	json_str, err := json.Marshal(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(json_str)
}

func CreateStream(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var data CreateStreamStruct
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	data_create := StreamData{
		Command:  "stream",
		Course:   data.Course,
		Room:     data.Room,
		Lecturer: data.Lecturer,
		Streamer: data.Streamer,
		Name:     data.Name,
		Stream:   "rtmp://live.edstrom.me/live",
	}

	var ansData StreamData
	res := SendStreamRequest("http://localhost:55994/", &data_create)
	err = json.Unmarshal(res, &ansData)
	if err != nil {
		panic(err)
	}

	// TODO
	token := ""
	res_data := CreateStreamResult{
		Key:     MakeKey(ansData.Id, token),
		Stream: "rtmp://live.edstrom.me/live",
	}

	res, err = json.Marshal(res_data)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/stream/create", CreateStream)

	http.Handle("/", r)
	server := &http.Server{
		Handler: r,
		Addr:    ":1339",
	}

	log.Fatal(server.ListenAndServe())
}
