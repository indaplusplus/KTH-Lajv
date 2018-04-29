package main

//User class, handles if they are authorized and contains info for comments

type user struct {
	id    string `json:"user"`
	token string `json:"token"`
}
