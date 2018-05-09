package main

//TODO: implement actual database interaction

var db map[string]string = map[string]string{"4T0k3n": "filip", "token": "usr"}

func loginToken(token string, user string) {
	db[token] = user
}

func logoutToken(token string) {
	delete(db, token)
}

func containsToken(token string) bool {
	_, has := db[token]
	return has
}

func getLoggedInUser(token string) (string, bool) {
	user, has := db[token]
	return user, has
}
