package main

import "fmt"

// func main() {
// 	fmt.Print("Hello World")
// }

type T struct {
	name  string
	value int
}

func main() {
	myvar := "Opeyemi"
	if myvar == "Ope" {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Found")
	}
}
