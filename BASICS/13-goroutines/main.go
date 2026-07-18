package main

import (
	"fmt"
	"sync"
	"time"
)

// This is Go's superpower and why companies hire Go engineers.
// A "goroutine" is an extremely lightweight thread. You start one with
// the "go" keyword. Channels let goroutines talk to each other safely.

func worker(id int, wg *sync.WaitGroup) {
	// defer runs this line when the function returns — here it tells the
	// WaitGroup "I'm done".
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond) // pretend to do work
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// --- 1) Running work concurrently with a WaitGroup ---
	// A WaitGroup lets main() wait until all goroutines finish.
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)        // "one more goroutine to wait for"
		go worker(i, &wg) // "go" launches it concurrently
	}
	wg.Wait() // block here until all 3 call wg.Done()
	fmt.Println("all workers finished")

	// --- 2) Channels: passing data between goroutines ---
	// A channel is a typed pipe. Send with ch <- v, receive with <-ch.
	ch := make(chan string)

	go func() {
		ch <- "hello from a goroutine" // send
	}()

	msg := <-ch // receive (waits until something is sent)
	fmt.Println("received:", msg)

	// --- 3) Collecting results from several goroutines via a channel ---
	results := make(chan int, 3) // buffered channel (holds up to 3)
	for i := 1; i <= 3; i++ {
		go func(n int) {
			results <- n * n // send the square back
		}(i)
	}
	sum := 0
	for i := 0; i < 3; i++ {
		sum += <-results
	}
	fmt.Println("sum of squares:", sum)
}
