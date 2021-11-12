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
	fetch3()
	// channels()
	// bufferedChannels()
	// directionalChannels()
	// controlFlow()
	// controlFlow2()
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
				m.Lock()
				cache[id] = b
				m.Unlock()
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond) // this bufffer time seems important
	}

	wg.Wait()
}

// select statements in order to not always get hits from db and cache (double reporting) - use a channel to pass info back and forth. 3rd goroutine to help make decisions about output
func fetch3() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}       // want a pointer bc we don't want to pass around copies to the mutex
	cacheCh := make(chan Book) // lets us know we received cache val - these channels are technically bidirectional - but if they're passed into functions with typed channels, that's okay, go knows what to do with them and will enforce that in that scope
	dbCh := make(chan Book)    // lets us know we received a db val
	for i := 0; i < 10; i++ {
		id := rnd.Intn(2) + 1

		// go keyword kicks off a go routine
		wg.Add(2)                                                              // or two calls to wg.Add(1) - this means the main goroutine knows we are waiting on 2 goroutines
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) { // send only channel
			if b, ok := queryCache(id, m); ok {
				ch <- b // send book to channel if we find a cache match
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// create one goroutine per query to generate a single response from the 2 executing goroutines above
		go func(cacheCh, dbCh <-chan Book) { // receive only channel
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh // this allows us to drain the dbCh of its message that we don't care about in this case so that we wait until we get the message from dbCh so we don't block the second goroutine
			case b := <-dbCh: // originally, messages in the dbCh never get drained (and they get put in each iteration above, regardless)
				fmt.Println("from db")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond) // this bufffer time seems important
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
			return b, true
		}
	}

	return Book{}, false
}
