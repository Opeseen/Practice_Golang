package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func retrieve(url string, wg *sync.WaitGroup) {
	// WaitGroup Counter-- when goroutine is finished
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(url)
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	// print the status code from the response
	fmt.Println(url, res.StatusCode, end)

}

// func main() {
// 	var wg sync.WaitGroup
// 	var urls = []string{"https://godoc.org", "https://www.packtpub.com", "https://kubernetes.io/"}
// 	for i := range urls {
// 		// WaitGroup Counter++ when new goroutine is called
// 		wg.Add(1)
// 		go retrieve(urls[i], &wg)
// 	}
// 	// Wait for the collection of goroutines to finish
// 	wg.Wait()
// }

// func main() {
// 	// Creating a buffered channel with a capacity of 3
// 	channel := make(chan int, 3)

// 	// Producer goroutine
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			channel <- i // Send number to channel
// 			fmt.Println("Produced:", i)
// 		}
// 		close(channel) // Close the channel when done producing
// 	}()

// 	// Consumer goroutine
// 	go func() {
// 		for num := range channel {
// 			fmt.Println("Consumed:", num)
// 			time.Sleep(time.Second) // Simulate time-consuming task
// 		}
// 	}()

// 	// Wait for a moment before exiting
// 	time.Sleep(6 * time.Second)
// }

// func main() {
// 	// Creating a buffered channel with a capacity of 2
// 	dataChannel := make(chan int, 2)

// 	go producer(dataChannel)
// 	go consumer(dataChannel)

// 	// Simulate work for 1 second before exiting
// 	time.Sleep(6 * time.Second)
// }

// producer sends data to the channel
func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // Non-blocking for the first 2 elements
		fmt.Println("Produced:", i)
	}
}

// consumer receives data from the channel
func consumer(ch chan int) {
	for {
		data := <-ch
		fmt.Println("Consumed:", data)
		time.Sleep(500 * time.Millisecond) // simulate time-consuming task
	}
}
