package main

import (
	"fmt"
	"net/http"

	"github.com/tzarick/go-learning/getting-started/starter/controllers"
	"github.com/tzarick/go-learning/getting-started/starter/foundation"
	"github.com/tzarick/go-learning/getting-started/starter/models"
)

func main() {
	fmt.Println("Ahoy gopher")

	foundation.Primitives()
	foundation.Collections()

	u := models.User{
		ID:        2,
		FirstName: "Pirate",
		LastName:  "Jack",
	}
	fmt.Println(u)

	foundation.Functions()

	// web server
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // nil uses the default mux
}
