package main

import (
	"fmt"
	"time"
)

// 01
func calculateSquare(number int) {
	time.Sleep(1 * time.Second)
	fmt.Println(number * number)
}

// func main() {
// 	start := time.Now()
// 	for i := 1; i <= 100; i++ {
// 		go calculateSquare(i)
// 	}
// 	elapsed := time.Since(start)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Function took: ", elapsed)
// }

// -------------------------------------------------------
// 02
func start() {
	go process()
	fmt.Println("In Start")
}

func process() {
	fmt.Println("In Process")
}

// func main() {
// 	go start()
// 	time.Sleep(1 * time.Second)
// }

// 03................................ Anonymous function

// func main() {
// 	go func() {
// 		fmt.Println("In anonymous method")
// 	}()
// 	time.Sleep(1 * time.Second)
// }

// 04.................................. GO routine Schedular
