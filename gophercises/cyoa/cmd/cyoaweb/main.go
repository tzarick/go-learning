package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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

	// tpl := template.Must(template.New("").Parse("Ahoy!")) // can pass a custom template in
	tpl := template.Must(template.New("").Parse(storyTemplate))

	h := cyoa.NewHandler(story, cyoa.WithTempate(tpl), cyoa.WithPathFunc(customPathFn))

	mux := http.NewServeMux()
	mux.Handle("/story/", h) // serve mux will return a 404 page if we don't go to /story

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

// custom options

func customPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
      <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li>
        <a href="/story/{{.Chapter}}">{{.Text}}</a>
      </li>
			{{end}}
    </ul>
  </body>
</html>`
