package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
