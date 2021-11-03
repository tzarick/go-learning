package foundation

import (
	"errors"
	"fmt"
)

func Functions() {
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")

	fmt.Println("Server started on port", port)
	fmt.Println("# of retires", numberOfRetries)

	if "something" == "nothing" {
		errors.New("something went wrong")
	}

	return port, nil
}
