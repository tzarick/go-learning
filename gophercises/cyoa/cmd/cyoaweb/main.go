package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tzarick/go-learning/gophercises/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the Choose Your Own Adventurer story")
	port := flag.Int("port", 3000, "the port to start the CYOA app on")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("error opening file [%s]: %v", *filename, err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
