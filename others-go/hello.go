package main

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

// func main() {
// 	var arr1 = [3]int{1, 2, 3}
// 	arr2 := [2]int{4, 5}
// 	var arr3 = [...]int{6, 7, 8}
// 	arr4 := [...]int{9, 10}
// 	arr5 := [5]int{1: 10, 2: 40}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
// 	fmt.Println(arr4)
// 	fmt.Println(arr4[1])
// 	arr4[1] = 30
// 	fmt.Println(arr4)
// 	fmt.Println(arr5)
// 	fmt.Println(len(arr5))
// }

// Creating Slice with data Type Values
// func main() {
// 	mySlice1 := []int{}
// 	mySlice2 := []string{"GO", "Slices", "Are", "Powerful"}
// 	fmt.Println(mySlice1)
// 	fmt.Println(len(mySlice1))
// 	fmt.Println(cap(mySlice1))

// 	fmt.Println(mySlice2)
// 	fmt.Println(len(mySlice2))
// 	fmt.Println(cap(mySlice2))
// }

// Create a Slice From an Array
// func main() {
// 	myArraay := [5]int{1, 2, 3, 4, 5}
// 	mySlice := myArraay[2:4]

// 	fmt.Printf("mySlice = %v\n", mySlice)
// 	fmt.Printf("length = %d\n", len(mySlice))
// 	fmt.Printf("Capacity = %d\n", cap(mySlice))
// }

// Create a Slice With The make() Function
// func main() {
// 	mySlice := make([]int, 5, 10)
// 	mySlice2 := make([]int, 5)
// 	mySlice2 = append(mySlice2, 1, 2, 3)
// 	fmt.Printf("mySlice = %v\n", mySlice)
// 	fmt.Printf("length = %d\n", len(mySlice))
// 	fmt.Printf("Capacity = %d\n", cap(mySlice))
// 	fmt.Println()
// 	fmt.Printf("mySlice2 = %v\n", mySlice2)
// 	fmt.Printf("length = %d\n", len(mySlice2))
// 	fmt.Printf("Capacity = %d\n", cap(mySlice2))
// 	fmt.Printf("mySlice2Append = %v\n", mySlice2)
// }

// func main() {
// 	mySlice := [3]int{1, 4, 6}
// 	fmt.Println(mySlice[1])
// 	mySlice[1] = 9
// 	fmt.Println(mySlice[1])
// }

// func main() {
// 	mySlice := []int{1, 2, 3}
// 	mySlice2 := []int{4, 5, 6}

// 	mySlice3 := append(mySlice, mySlice2...)
// 	fmt.Println(mySlice3)
// 	fmt.Printf("Capacity %d\n", cap(mySlice3))
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
// 	// Original slice
// 	fmt.Printf("numbers = %v\n", numbers)
// 	fmt.Printf("length = %d\n", len(numbers))
// 	fmt.Printf("capacity = %d\n", cap(numbers))
// 	fmt.Println()

// 	// Create copy with only needed numbers
// 	neededNumbers := numbers[:len(numbers)-10]
// 	neededNumbers2 := numbers[5:]
// 	numberCopy := make([]int, len(neededNumbers))
// 	copy(numberCopy, neededNumbers)
// 	fmt.Println(neededNumbers)
// 	fmt.Println(neededNumbers2)
// 	fmt.Println(numberCopy)
// }
