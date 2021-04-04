package app

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/github/login", GithubCallbackHandler)

	return r
}
