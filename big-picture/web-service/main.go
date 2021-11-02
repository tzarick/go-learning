package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name}
		enc := json.NewEncoder(rw) // encoder to encode the map to json. Anything we encode with encoder will encode it to the writer we pass in
		enc.Encode(m)

		rw.Write([]byte("Hello " + name))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
