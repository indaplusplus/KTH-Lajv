package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"encoding/json"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "./kthlive.db")
	if err != nil {
		log.Fatal(err)
	}
	if !tablesExist(db) {
		createTables(db)
	}
	http.Handle("/", httpHandler{db})
	log.Fatal(http.ListenAndServe(":55994", nil))
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
		replyToUser VARCHAR(50),
		replyToTime TIMESTAMP,
		PRIMARY KEY(user, time),
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
	_, err = db.Exec("CREATE TABLE Logins (token VARCHAR(50), user VARCHAR(50));")
	if err != nil {
		log.Fatal(err)
	}
}

type httpHandler struct {
	db *sql.DB
}

func (handler httpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	var data jsonData
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Print(err)
		return
	}
	switch data.Command {
	case "stream":
		res, err := json.Marshal(stream(data, handler.db))
		if err != nil {
			log.Print(err)
			return
		}
		writer.Write(res)
	case "stop-stream":
		stopStream(data, handler.db)
	case "find":
		res, err := json.Marshal(find(data, handler.db))
		if err != nil {
			log.Print(err)
			return
		}
		writer.Write(res)
	case "watch":
		res, err := json.Marshal(watch(data, handler.db))
		if err != nil {
			log.Print(err)
			return
		}
		writer.Write(res)
	case "chat":
		chat(data, handler.db)
	case "comment":
		comment(data, handler.db)
	case "get-chat":
		res, err := json.Marshal(getChat(data, handler.db))
		if err != nil {
			log.Print(err)
			return
		}
		writer.Write(res)
	case "get-comments":
		res, err := json.Marshal(getComments(data, handler.db))
		if err != nil {
			log.Print(err)
			return
		}
		writer.Write(res)
	case "upvote-comment":
		upvoteComment(data, handler.db)
	case "delete-comment":
		deleteComment(data, handler.db)
	case "login":
		login(data, handler.db)
	case "loggedin":
		res, err := json.Marshal(loggedIn(data, handler.db))
		if err != nil {
			log.Print(err)
			return
		}
		writer.Write(res)
	case "logout":
		logout(data, handler.db)
	}
}

type jsonData struct {
	Command     string         `json:"command"`
	Course      string         `json:"course"`
	Room        string         `json:"room"`
	Lecturer    string         `json:"lecturer"`
	Streamer    string         `json:"streamer"`
	Name        string         `json:"name"`
	Date        string         `json:"date"`
	Vod         string         `json:"vod"`
	Stream      string         `json:"stream"`
	Hls         string         `json:"hls"`
	Id          int            `json:"id"`
	Ids         [] int         `json:"ids"`
	User        string         `json:"user"`
	Time        string         `json:"time"`
	Text        string         `json:"text"`
	ReplyToUser string         `json:"replyToUser"`
	ReplyToTime string         `json:"replyToTime"`
	Chat        [] messageData `json:"chat"`
	Comments    [] messageData `json:"comments"`
	Token       string         `json:"token"`
}

type messageData struct {
	User        string `json:"user"`
	Time        string `json:"time"`
	Text        string `json:"text"`
	Upvotes     int    `json:"upvotes"`
	ReplyToUser string `json:"replyToUser"`
	ReplyToTime string `json:"replyToTime"`
}

func stream(data jsonData, db *sql.DB) (returns jsonData) {
	res, err := db.Query("SELECT COUNT(*) FROM Streams;")
	if err != nil {
		log.Print(err)
		return
	}
	var count int
	res.Next()
	res.Scan(&count)
	res.Close()
	_, err = db.Exec("INSERT INTO Streams VALUES(?, ?, ?, ?, ?, ?, CURRENT_DATE, ?, ?, ?);",
		count, data.Course, data.Room, data.Lecturer, data.Streamer, data.Name, "", data.Stream, data.Hls)
	if err != nil {
		log.Print(err)
		return
	}
	returns.Id = count
	return
}

func stopStream(data jsonData, db *sql.DB) {
	_, err := db.Exec("UPDATE Streams SET vod = ?, stream = ?, hls = ? WHERE id = ?",
		data.Vod, "", "", data.Id)
	if err != nil {
		log.Print(err)
	}
	return
}

func find(data jsonData, db *sql.DB) (returns jsonData) {
	query := "SELECT id FROM Streams WHERE "
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
	res, err := db.Query(query[:len(query)-5]+";", values...)
	if err != nil {
		log.Print(err)
		return
	}
	for res.Next() {
		var id int
		res.Scan(&id)
		returns.Ids = append(returns.Ids, id)
	}
	res.Close()
	return
}

func watch(data jsonData, db *sql.DB) (returns jsonData) {
	res, err := db.Query("SELECT vod, stream, hls FROM streams WHERE id = ?", data.Id)
	if err != nil {
		log.Print(err)
		return
	}
	var vod, stream, hls string
	res.Next()
	res.Scan(&vod, &stream, &hls)
	res.Close()
	returns.Vod = vod
	returns.Stream = stream
	returns.Hls = hls
	return
}

func chat(data jsonData, db *sql.DB) (returns jsonData) {
	_, err := db.Exec("INSERT INTO ChatMessages VALUES(?, ?, CURRENT_TIME, ?);",
		data.Id, data.User, data.Text)
	if err != nil {
		log.Print(err)
	}
	return
}

func comment(data jsonData, db *sql.DB) (returns jsonData) {
	_, err := db.Exec("INSERT INTO Comments VALUES(?, ?, CURRENT_TIMESTAMP, ?, 0, ?, ?);",
		data.Id, data.User, data.Text, data.ReplyToUser, data.ReplyToTime)
	if err != nil {
		log.Print(err)
	}
	return
}

func getChat(data jsonData, db *sql.DB) (returns jsonData) {
	res, err := db.Query("SELECT user, time, text FROM ChatMessages WHERE id = ?;",
		data.Id)
	if err != nil {
		log.Print(err)
	}
	for res.Next() {
		var user, time, text string
		res.Scan(&user, &time, &text)
		returns.Chat = append(returns.Chat, messageData{user, time, text, 0, "", ""})
	}
	res.Close()
	return
}

func getComments(data jsonData, db *sql.DB) (returns jsonData) {
	res, err := db.Query("SELECT user, time, text, upvotes, replyToUser, replyToTime FROM Comments WHERE id = ?;",
		data.Id)
	if err != nil {
		log.Print(err)
	}
	for res.Next() {
		var user, time, text, replyToUser, replyToTime string
		var upvotes int
		res.Scan(&user, &time, &text, &upvotes, &replyToUser, &replyToTime)
		returns.Comments = append(returns.Comments, messageData{user, time, text, upvotes, replyToUser, replyToTime})
	}
	res.Close()
	return
}

func upvoteComment(data jsonData, db *sql.DB) (returns jsonData) {
	_, err := db.Exec("UPDATE Comments SET upvotes = upvotes + 1 WHERE user = ? AND time = ?;", data.User, data.Time)
	if err != nil {
		log.Print(err)
	}
	return
}

func deleteComment(data jsonData, db *sql.DB) (returns jsonData) {
	_, err := db.Exec("DELETE FROM Comments WHERE user = ? AND time = ?;", data.User, data.Time)
	if err != nil {
		log.Print(err)
	}
	return
}

func login(data jsonData, db *sql.DB) (returns jsonData) {
	_, err := db.Exec("INSERT INTO Logins VALUES(?, ?);", data.Token, data.User)
	if err != nil {
		log.Print(err)
	}
	return
}

func loggedIn(data jsonData, db *sql.DB) (returns jsonData) {
	res, err := db.Query("SELECT user FROM Logins WHERE token = ?;", data.Token)
	if err != nil {
		log.Print(err)
	}
	var user string
	res.Next()
	res.Scan(&user)
	res.Close()
	returns.User = user
	return
}

func logout(data jsonData, db *sql.DB) (returns jsonData) {
	_, err := db.Exec("DELETE FROM Logins WHERE token = ?;", data.Token)
	if err != nil {
		log.Print(err)
	}
	return
}
