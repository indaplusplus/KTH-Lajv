package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"encoding/json"
	"fmt"
)

func main() {
	db, err := sql.Open("sqlite3", "./kthlive.db")
	if err != nil {
		log.Fatal(err)
	}
	if !tablesExist(db) {
		createTables(db)
	}
	db.Close()
}

type id struct {
	Id int  `json:"id"`
}

type ids struct {
	Ids[] int `json:"ids"`
}

func tablesExist(db *sql.DB) (exists bool) {
	res, err := db.Query("SELECT EXISTS (SELECT * FROM sqlite_master WHERE TYPE = 'table');")
	if err != nil {
		log.Fatal(err)
	}
	res.Next()
	res.Scan(&exists)
	res.Close()
	return
}

func createTables(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE Streams (
		id INTEGER PRIMARY KEY,
		course VARCHAR(10),
		room VARCHAR(10),
		lecturer VARCHAR(50),
		streamer VARCHAR(50),
		name VARCHAR(50),
		date DATE,
		vod VARCHAR(100),
		stream VARCHAR(100),
		hls VARCHAR(100));`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE Comments (
		id INTEGER,
		user VARCHAR(50),
		time TIMESTAMP,
		text VARCHAR(500),
		upvotes INTEGER,
		FOREIGN KEY(id) REFERENCES Streams(id));`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE ChatMessages (
		id INTEGER,
		user VARCHAR(50),
		time TIME,
		text VARCHAR(500),
		FOREIGN KEY(id) REFERENCES Streams(id));`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE Logins (token VARCHAR(50));")
	if err != nil {
		log.Fatal(err)
	}
}

type streamData struct {
	Course string `json:"course"`
	Room string `json:"room"`
	Lecturer string `json:"lecturer"`
	Streamer string `json:"streamer"`
	Name string `json:"name"`
	Stream string `json:"stream"`
	Hls string `json:"hls"`
}

func stream(params []byte, db *sql.DB) (returns []byte, err error) {
	var data streamData
	err = json.Unmarshal(params, &data)
	if err != nil {
		return
	}
	res, err := db.Query("SELECT COUNT(*) FROM streams;")
	if err != nil {
		return
	}
	var count int
	res.Next()
	res.Scan(&count)
	res.Close()
	_, err = db.Exec("INSERT INTO streams VALUES(?, ?, ?, ?, ?, ?, CURRENT_DATE, ?, ?, ?);",
		count, data.Course, data.Room, data.Lecturer, data.Streamer, data.Name, "", data.Stream, data.Hls)
	if err != nil {
		return
	}
	returns, err = json.Marshal(id{count})
	return
}

type findData struct {
	Course string `json:"course"`
	Room string `json:"room"`
	Lecturer string `json:"lecturer"`
	Streamer string `json:"streamer"`
	Name string `json:"name"`
	Date string `json:"date"`
}

func find(params []byte, db *sql.DB) (returns []byte, err error) {
	var data findData
	err = json.Unmarshal(params, &data)
	if err != nil {
		return
	}
	query := "SELECT id FROM streams WHERE "
	var values []interface{}
	if len(data.Course) > 0 {
		query += "course = ? AND "
		values = append(values, data.Course)
	}
	if len(data.Room) > 0 {
		query += "room = ? AND "
		values = append(values, data.Room)
	}
	if len(data.Lecturer) > 0 {
		query += "lecturer = ? AND "
		values = append(values, data.Lecturer)
	}
	if len(data.Streamer) > 0 {
		query += "streamer = ? AND "
		values = append(values, data.Streamer)
	}
	if len(data.Name) > 0 {
		query += "name = ? AND "
		values = append(values, data.Name)
	}
	if len(data.Date) > 0 {
		query += "date = ? AND "
		values = append(values, data.Date)
	}
	res, err := db.Query(query[:len(query)-5] + ";", values...)
	if err != nil {
		return
	}
	var ids ids
	for res.Next() {
		fmt.Println("here")
		var id int
		res.Scan(&id)
		ids.Ids = append(ids.Ids, id)
	}
	res.Close()
	returns, err = json.Marshal(ids)
	return
}

type watchData struct {
	Vod string `json:"vod"`
	Stream string `json:"stream"`
	Hls string `json:"hls"`
}

func watch(params []byte, db *sql.DB) (returns []byte, err error) {
	var id id
	err = json.Unmarshal(params, &id)
	if err != nil {
		return
	}
	res, err := db.Query("SELECT vod, stream, hls FROM streams WHERE id = ?", id.Id)
	if err != nil {
		return
	}
	var vod, stream, hls string
	res.Next()
	res.Scan(&vod, &stream, &hls)
	res.Close()
	returns, err = json.Marshal(watchData{vod, stream, hls})
	return
}