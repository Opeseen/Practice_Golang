package main

import "fmt"

// func main() {
// 	fmt.Print("Hello World")
// }

type T struct {
	name  string
	value int
}

// func main() {
// 	myvar := "Opeyemi"
// 	if myvar == "Ope" {
// 		fmt.Println("Not Found")
// 	} else {
// 		fmt.Println("Found")
// 	}
// }

// func main() {
// 	var b int
// 	var a string
// 	var c bool

// 	fmt.Println(b, a, c)
// }

// func main() {
// 	name := "Opeyemi"
// 	fmt.Println("Old name is", name)
// 	name = "Mike"
// 	fmt.Println("New name is ", name)
// }

var name string = "Opeyemi"
var name2 string = "Mike"

// func main() {
// 	name = "Ope"
// 	fmt.Println((name))
// 	name2 = "Mikel"
// 	fmt.Println(name2)
// }

// func main() {
// 	var a, b, c, d float64 = 1, 2, 3, 4
// 	var1, var2 := 5, 6
// 	fmt.Println(a, b, c, d, var1, var2)
// }

// func main() {
// 	var (
// 		a     int    = 5
// 		name  string = "Ope"
// 		age   int    = 25
// 		state string = "Lagos"
// 	)

// 	fmt.Println(a, age, name, state)
// }

// func main() {
// 	const (
// 		name  = "Ope"
// 		age   = 25
// 		state = "Lagos"
// 		code  = 234
// 	)
// 	fmt.Println(name, age, state)
// 	fmt.Print(age, code)
// }

// func main() {
// 	var i string = "Ope"
// 	const j int = 26
// 	const b2 = true
// 	const num uint = 43

// 	fmt.Printf("I haa a value of %v and type %T", i, i)
// 	fmt.Println()
// 	fmt.Printf("J has a value of %v and type %T", j, j)
// 	fmt.Println()
// 	fmt.Printf("b2 has a value of %v and a tye %T", b2, b2)
// 	fmt.Println()
// 	fmt.Println(num)
// }

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [2]int{4, 5}
	var arr3 = [...]int{6, 7, 8}
	arr4 := [...]int{9, 10}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
