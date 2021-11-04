package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/tzarick/go-learning/getting-started/starter/models"
)

// some object oriented techniques. userController is like a class?

type userController struct {
	userIDPattern *regexp.Regexp
}

// bind method to an object. uc is like `this` - this is an instance method
// servehttp decides where to pass the request
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			println("")
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			println("put")
		case http.MethodDelete:
			println("delete")
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
	w.Write([]byte("Hello hello from userController"))
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJson(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {

}

// constructor function -> convention: start with `new` - often return a pointer to it
func newUserController() *userController {
	return &userController{ // can immediately take the address of a struct literal. This is a local variable, which is in the scope of this function. However, Go recognizes that we are returning this memory location and will not free it after returning
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
