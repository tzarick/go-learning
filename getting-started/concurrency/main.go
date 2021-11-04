package main

import (
	"fmt"
	"math/rand"
	"time"
)

// create a cache and a "DB" of books

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fetch1()
}

func fetch1() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(2) + 1

		// go keyword kicks off a go routine
		go func(id int) { // anonymous function for immediate execution
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from db")
				fmt.Println(b)
			}
		}(id)
		// fmt.Printf("Book not found with id '%v'\n\n", id)
		time.Sleep(150 * time.Millisecond) // sleep call here gives some time for the go routines to complete (we know the longest routine will be 100ms + a little so they all should have completed before we exit the program - otherwise, we will kick off all the goroutines and immediately exit before getting any results)
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond) // simulate slowness to access the db
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}
