package main

//TODO: implement actual database interaction

var db map[string]bool = map[string]bool{}

func loginToken(token string) {
  db[token] = true
}

func logoutToken(token string) {
  delete(db, token)
}

func containsToken(token string) bool {
  return db[token]
}
