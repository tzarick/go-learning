package main

import (
	"fmt"
	"sync"
)

func channels() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		fmt.Println(<-ch) // receiving a msg from the channel. Receive data from the channel
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		ch <- 42 // pass a message into the channel. Send data to the channel
		// ch <- 72 // adding this line will cause another deadlock because there isn't ANOTHER receiver on the other end. The sender/receiver count doesn't match up
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func tryToDoTheSameThingWithoutGoroutines() {
	ch := make(chan int)
	// we can't receive data from a channel until there is data being sent through it (otherwise we will block here until there is something sent through) -> deadlock!
	fmt.Println(<-ch)
	ch <- 42

	// we can't send data through a channel until there is a receiver listening for data on the other side (otherwise we will block here until there is) -> deadlock!
	ch <- 42
	fmt.Println(<-ch)
}

func bufferedChannels() {
	wg := &sync.WaitGroup{}
	// internal capacity of an UNBUFFERED channel is zero aka you have to have a matching number of senders / receivers
	ch := make(chan int, 1) // in this case we can have 1 msg sitting in the buffer (don't need equal tx / rx count)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		fmt.Println(<-ch) // receiving a msg from the channel. Receive data from the channel
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		ch <- 42 // pass a message into the channel. Send data to the channel
		ch <- 72
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

// often we only want a goroutine to either send or receive
func directionalChannels() {
	wg := &sync.WaitGroup{}
	// internal capacity of an UNBUFFERED channel is zero aka you have to have a matching number of senders / receivers
	ch := make(chan int, 1) // in this case we can have 1 msg sitting in the buffer (don't need equal tx / rx count)

	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		fmt.Println(<-ch) // receiving a msg from the channel. Receive data from the channel
		wg.Done()
	}(ch, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		ch <- 42 // pass a message into the channel. Send data to the channel
		// close(ch)
		ch <- 72
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func controlFlow() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // in this case we can have 1 msg sitting in the buffer (don't need equal tx / rx count)

	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly

		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok) // receiving a msg from the channel. Receive data from the channel
		} else {
			fmt.Println("closed")
		}

		wg.Done()
	}(ch, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) { // best to pass the vars in instead of using them directly
		ch <- 2000
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
