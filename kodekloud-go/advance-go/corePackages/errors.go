package main

import (
	"errors"
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		fmt.Println(errors.New("Only odd numbers allowed"))
		return fmt.Errorf("Only odd numbers allowed, got: %d", i)
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Process ran successfully")
}

// func main() {
// 	err := process(3)
// 	checkError(err)
// 	err = process(2)
// 	checkError(err)
// }
