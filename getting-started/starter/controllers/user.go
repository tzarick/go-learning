package controllers

import (
	"net/http"
	"regexp"
)

// some object oriented techniques. userController is like a class?

type userController struct {
	userIDPattern *regexp.Regexp
}

// bind method to an object. uc is like `this` - this is an instance method
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello hello from userController"))
}

// constructor function -> convention: start with `new` - often return a pointer to it
func newUserController() *userController {
	return &userController{ // can immediately take the address of a struct literal. This is a local variable, which is in the scope of this function. However, Go recognizes that we are returning this memory location and will not free it after returning
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
