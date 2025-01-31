package main

import (
	"fmt"
	"time"
)

func calculateSquare(number int) {
	time.Sleep(1 * time.Second)
	fmt.Println(number * number)
}

func main() {
	start := time.Now()
	for i := 1; i <= 100; i++ {
		go calculateSquare(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took: ", elapsed)
}
