package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	linkChannel := make(chan string)
// 	go func() {
// 		links := []string{"https://godoc.org", "https://www.packtpub.com", "https://kubernetes.io/"}
// 		for _, link := range links {
// 			linkChannel <- link
// 		}
// 	}()
// 	link1 := <-linkChannel
// 	link2 := <-linkChannel
// 	link3 := <-linkChannel

// 	fmt.Println(link1)
// 	fmt.Println(link2)
// 	fmt.Println(link3)
// }

func ping(c chan string) {
	links := []string{"https://godoc.org", "https://www.packtpub.com", "https://kubernetes.io/"}
	for _, link := range links {
		c <- link
	}
	close(c)
}

// func main() {
// 	linkChannel := make(chan string)
// 	go ping(linkChannel)

// 	link1 := <-linkChannel
// 	link2 := <-linkChannel
// 	link3 := <-linkChannel

// 	fmt.Println(link1)
// 	fmt.Println(link2)
// 	fmt.Println(link3)

// 	for {
// 		_, ok := <-linkChannel
// 		if ok {
// 			fmt.Println("Channel is Open")
// 		} else {
// 			fmt.Println("Channel is Closed")
// 			break
// 		}
// 	}
// }

// .........................Buffered Channel
func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	go buy(ch, wg)
	fmt.Println("Sent all data to the channel")
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	fmt.Println("Received data", <-ch)
	fmt.Println("Received data", <-ch)
	fmt.Println("Received data", <-ch)
	wg.Done()
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	ch := make(chan int, 3)
// 	go sell(ch, &wg)
// 	wg.Wait()
// }

// ................Closed channel
// func main() {
// 	ch := make(chan int, 10)
// 	ch <- 10
// 	ch <- 11
// 	val, ok := <-ch
// 	fmt.Println(val, ok)
// 	close(ch)
// 	val, ok = <-ch
// 	fmt.Println(val, ok)
// 	val, ok = <-ch
// 	fmt.Println(val, ok)
// }

// .............................Panic
// func main() {
// 	ch := make(chan int, 10)
// 	ch <- 10
// 	ch <- 11
// 	val, ok := <-ch
// 	fmt.Println(val, ok)
// 	close(ch)
// 	ch <- 12
// }

// func main() {
// 	ch := make(chan int, 5)
// 	ch <- 10
// 	ch <- 20
// 	ch <- 30
// 	close(ch)

// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

// ...................................Select Statement
func goOne(ch1 chan string) {
	ch1 <- "Channel-1"
}

func goTwo(ch2 chan string) {
	ch2 <- "Channel-2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
		break
	case val2 := <-ch2:
		fmt.Println(val2)
		break
	default:
		fmt.Println("Executed default")
	}
	time.Sleep(3 * time.Second)
}
