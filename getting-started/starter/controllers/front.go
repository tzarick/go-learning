package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc) // uc implements interface http.Handler (because it has a serveHTTP method) which is the second arg to this method
	http.Handle("/users/", *uc)
}
