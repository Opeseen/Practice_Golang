package main

import (
	"fmt"
	"sync"
	"time"
)

// ........................concurrency leak

// This will throw an error
func leak(wg *sync.WaitGroup) {
	ch := make(chan string)

	go func() {
		val := <-ch
		fmt.Println("Value received from channel: ", val)
		wg.Done()
	}()

	fmt.Println("Exiting leak method")
	wg.Done()
}

// func main() {
// 	var wg = sync.WaitGroup{}
// 	wg.Add(2)
// 	go leak(&wg)
// 	wg.Wait()
// }

// func main() {
// 	var wg = sync.WaitGroup{}
// 	wg.Add(10)

// 	for i := 0; i <= 10; i++ {
// 		go func(i int) {
// 			fmt.Println(i)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()

// 	fmt.Println("Done")
// }

func sendValue(ch1 chan int) {
	time.Sleep(3 * time.Second)
	ch1 <- 10
}

// func main() {
// 	ch1 := make(chan int)
// 	go sendValue(ch1)

// 	select {
// 	case msg := <-ch1:
// 		fmt.Println(msg)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Select timeout")
// 	}
// }
