package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Sum Functions
func sumNumber(a int, b int) int {
	return a + b
}

// func main() {
// 	sum := sumNumber(10, 20)

// 	fmt.Println(sum)
// }

// Variadic Functions
func addMultiplesNumber(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// func main() {
// 	result := addMultiplesNumber(10, 20, 30)

// 	fmt.Println(result)
// 	fmt.Println((addMultiplesNumber()))
// 	fmt.Println((addMultiplesNumber(450, 60)))
// 	fmt.Println(addMultiplesNumber(30, 0))
// }

func concatString1(words ...string) string {
	sum := "Completed"
	for _, y := range words {
		fmt.Println((y))
	}
	return sum
}

// func main() {
// 	fmt.Println((concatString("Hello", "How")))
// }

func concatString2(words ...string) string {
	sum := "Completed"
	for x, y := range words {
		fmt.Println(strconv.Itoa(x), y)
	}
	return sum
}

// func main() {
// 	fmt.Println((concatString("Hello", "How", "Are", "You")))
// }

func StudentGrades(student string, grades ...string) {
	fmt.Println("Hello", student, " see your grades below")

	for _, score := range grades {
		fmt.Printf("%s ", score)
	}
}

// func main() {
// 	StudentGrades("Mike", "A+", "B+", "C")
// }

func f() (int, int) {
	return 42 + 43, 45
}

// func main() {
// 	a, b := f()
// 	x, _ := f()
// 	_, y := f()

// 	fmt.Println(a, b)
// 	fmt.Println(x)
// 	fmt.Println((y))
// }

func calcSquare(numbers []int) []int {
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
	}
	return squares

}

// func main() {
// 	nums := [3]int{10, 20, 15}
// 	fmt.Println(calcSquare(nums[:]))
// }

func calcSquare1(numbers []int) []int {
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
	}
	return squares
}

// func main() {
// 	nums := [3]int{10, 20, 15}
// 	fmt.Println(calcSquare(nums[:]))
// }

func calcSquare2(numbers []int) ([]int, bool) {
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
	}
	return squares, false
}

// func main() {
// 	nums := [3]int{10, 20, 15}
// 	fmt.Println(calcSquare(nums[:]))
// }

func printStrings1(s string, names ...string) {
	fmt.Println(s)
	for _, value := range names {
		fmt.Printf("%s, ", value)
	}
}

// func main() {
// 	printStrings("Hey there", "Joe", "Monica", "Gunther")
// }

func printStrings2(names string, numbers []int) (names_c string, num_c []int) {
	names_c = names
	num_c = numbers
	for _, values := range numbers {
		fmt.Println(names, " ", values)
	}
	return
}

// func main() {
// 	nums := [3]int{10, 20, 15}
// 	result1, result2 := printStrings("Mike", nums[:])
// 	fmt.Println(result1, " ", result2)
// }

// Recursive Function

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	fmt.Println(num, "*", factorial(num-1))
	return num * factorial(num-1)
}

// func main() {
// 	n := 5
// 	fmt.Println("Factorial of", n, "is", factorial(n))
// }

// Anonymous Function
// func main() {
// 	anonymous := func(x int, y int) int {
// 		return x * y
// 	}
// 	fmt.Printf("%T \n", anonymous)
// 	fmt.Println(anonymous(10, 20))
// }

// func main() {
// 	anonymous := func(x int, y int) int {
// 		return x * y
// 	}(20, 30)

// 	fmt.Printf("%T \n", anonymous)
// 	fmt.Println(anonymous)
// }

// func main() {
// 	tested := func(text string) string {
// 		return strings.ToLower(text)
// 	}("RACHAEL")
// 	fmt.Printf("%T \n", tested)
// 	fmt.Println(tested)
// }

//-- Higher Order Function

func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

// func main() {
// 	var query int
// 	var radius float64

// 	fmt.Print("Enter the radius of the circles: ")
// 	fmt.Scanf("%f", &radius)
// 	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
// 	fmt.Scanf("%d", &query)

// 	if query == 1 {
// 		fmt.Println("Result", calcArea(radius))
// 	} else if query == 2 {
// 		fmt.Println("Result", calcPerimeter(radius))
// 	} else if query == 3 {
// 		fmt.Println("Result", calcDiameter(radius))
// 	} else {
// 		fmt.Println("Invalid Selection", query)
// 	}
// }

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("result: ", result)
	fmt.Println("Thank You!")
}

func getFunction(query int) func(r float64) float64 {

	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

// func main() {
// 	var query int
// 	var radius float64

// 	fmt.Print("Enter the radius of the circles: ")
// 	fmt.Scanf("%f", &radius)
// 	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
// 	fmt.Scanf("%d", &query)

// 	printResult(radius, getFunction(query))
// }

func addHundred(x int) int {
	return x + 100
}
func partialSum(x ...int) func() {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return func() {
		fmt.Println(addHundred(sum))
	}
}

// func main() {
// 	partial := partialSum(1, 2, 3, 4, 5)
// 	partial()
// }

// Defer Statement
func printName(name string) {
	fmt.Println(name)
}

func printNumber(number int) {
	fmt.Println(number)
}

func printAddress(address string) {
	fmt.Println(address)
}

// func main() {
// 	defer printName("Horpeyemi")
// 	printNumber(29)
// 	printAddress("13, Mulero Street")
// }

func getString(str string) (string, string) {
	return strings.ToLower(str), strings.ToUpper(str)
}

// func main() {
// 	_, lower := getString("BROWSER")
// 	fmt.Println(lower)
// }

func greetings() (x, y string) {
	x = "hello "
	y = "world"
	return
}

func main() {
	fmt.Print(greetings())
}
