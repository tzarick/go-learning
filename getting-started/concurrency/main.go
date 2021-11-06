package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// create a cache and a "DB" of books

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// fetch1()
	// fetch2()
	// channels()
	// bufferedChannels()
	// directionalChannels()
	controlFlow()
	// tryToDoTheSameThingWithoutGoroutines()
}

func fetch1() {
	// for i := 0; i < 10; i++ {
	// 	id := rnd.Intn(2) + 1

	// 	// go keyword kicks off a go routine
	// 	go func(id int) { // anonymous function for immediate execution
	// 		if b, ok := queryCache(id); ok {
	// 			fmt.Println("from cache")
	// 			fmt.Println(b)
	// 		}
	// 	}(id)
	// 	go func(id int) {
	// 		if b, ok := queryDatabase(id); ok {
	// 			fmt.Println("from db")
	// 			fmt.Println(b)
	// 		}
	// 	}(id)
	// 	// fmt.Printf("Book not found with id '%v'\n\n", id)
	// 	time.Sleep(150 * time.Millisecond) // sleep call here gives some time for the go routines to complete (we know the longest routine will be 100ms + a little so they all should have completed before we exit the program - otherwise, we will kick off all the goroutines and immediately exit before getting any results)
	// }
}

// introduce WaitGroup
func fetch2() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{} // want a pointer bc we don't want to pass around copies to the mutex
	for i := 0; i < 10; i++ {
		id := rnd.Intn(2) + 1

		// go keyword kicks off a go routine
		wg.Add(2)                                              // or two calls to wg.Add(1) - this means the main goroutine knows we are waiting on 2 goroutines
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) { // anonymous function for immediate execution
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("from db")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
	}

	wg.Wait()
}

//// helpers

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock() // we want to use read/write mutex here instead because there will be many more reads from the cache than writing (having multiple reads is okay as long as there is no concurrent writing)
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond) // simulate slowness to access the db
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return Book{}, false
}
