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

// func main() {
// 	fmt.Print(greetings())
// }

// Pointers => Referencing & Dereferencing
// func main() {
// 	i := 10
// 	fmt.Printf("%T %v \n", &i, &i)
// 	fmt.Printf("%T %v \n", *(&i), *(&i))
// }

// func main() {
// 	s := "hello"
// 	var b *string = &s
// 	fmt.Println(b)
// 	var a = &s
// 	fmt.Println(a)
// 	c := &s
// 	fmt.Println(c)
// }

// Dereferencing Pointers
// func main() {
// 	s := "hello"
// 	fmt.Printf("%T %v \n", s, s)
// 	ps := &s
// 	*ps = "world"
// 	fmt.Printf("%T %v \n", s, s)
// }

// func main() {
// 	y := [3]int{10, 20, 30}
// 	fmt.Printf("%v \n", y)
// 	(*&y)[0] = 100
// 	fmt.Printf("%v \n", y)
// }

// func main() {
// 	var y int
// 	var ptr *int = &y

// 	*ptr = 0
// 	fmt.Println(y)

// 	*ptr += 5
// 	fmt.Println(y)

// }

// func main() {
// 	s := "hello"
// 	var ptr *string = &s
// 	fmt.Println(s)
// 	*ptr += strings.ToUpper(s)
// 	fmt.Println(s)
// }

// func modify(numbers ...int) {
// 	for i := range numbers {
// 		numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := []int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr...)
// 	fmt.Println(arr)
// }

// func modify(numbers [3]int) {
// 	for i, x := range numbers {
// 		numbers[i] -= 5
// 		fmt.Println(i, x)
// 	}
// }
// func main() {
// 	arr := [3]int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr)
// 	fmt.Println(arr)
// }

// func modify(numbers *[3]int) {
// 	for i := range numbers {
// 		numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := [3]int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr)
// 	fmt.Println(arr)
// }

// func modify(s *string) {
// 	*s = strings.ToUpper(*s)
// }
// func main() {
// 	s := "hello"
// 	fmt.Println(s)
// 	modify(&s)
// 	fmt.Println(s)
// }

func modify(s map[string]int) {
	s["A"] = 100
}

// func main() {
// 	ascii_codes := map[string]int{}
// 	ascii_codes["A"] = 65
// 	fmt.Println(ascii_codes)
// 	modify(ascii_codes)
// 	fmt.Println(ascii_codes)
// }

// STRUCT

type Student struct {
	name   string
	rollNo int
	mark   [3]int
	grades map[string]int
}

// func main() {
// 	var s Student
//
// fmt.Printf("%+v", s)
// }

// func main() {
// 	st := new(Student)
// 	fmt.Printf("%+v", st)
// }

// func main() {
// 	st := Student{
// 		name:   "Joe",
// 		rollNo: 15,
// 	}
// 	fmt.Printf("%+v", st)
// }

// func main() {
// 	var st Student
// 	st.name = "Opeyemi"
// 	st.rollNo = 12
// 	st.mark = [3]int{10, 20, 15}
// 	fmt.Printf("%+v", st)
// }

// Passing Struct to function
type Circles struct {
	x      int
	y      int
	radius float64
	area   float64
}

func calculateArea(c Circles) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	c.area = area
}

// func main() {
// 	c := Circles{x: 5, y: 5, radius: 5, area: 0}
// 	fmt.Printf("%+v \n", c)
// 	calculateArea(c)
// 	fmt.Printf("%+v \n", c)
// }

func calculateArea2(c *Circles) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	(*c).area = area
}

// func main() {
// 	c := Circles{x: 5, y: 5, radius: 5, area: 0}
// 	fmt.Printf("%+v \n", c)
// 	calculateArea2(&c)
// 	fmt.Printf("%+v \n", c)
// }

// Comparing Struct
type s1 struct {
	x int
}
type s2 struct {
	x int
}

// func main() {
// 	c := s1{x: 5}
// 	c1 := s2{x: 5}
// 	if(c == c1){
// 		fmt.Println('Yes')
// 	}
// }

// func main() {
// 	c := s1{x: 5}
// 	c1 := s1{x: 6}
// 	c2 := s1{x: 5}

// 	if c != c1 {
// 		fmt.Println("c and c1 have different values")
// 	}
// 	if c == c2 {
// 		fmt.Println("c is same as c2")
// 	}
// }

// Methods
type Circle2 struct {
	radius float64
	area   float64
}

type Circle3 struct {
	radius float64
	area   float64
}

func (c *Circle2) calcArea2() {
	c.area = 3.14 * c.radius * c.radius
}

// func main() {
// 	c := Circle2{radius: 5}
// 	c.calcArea2()
// 	fmt.Printf("%+v", c)
// }

func (c Circle3) calcArea3() {
	c.area = 3.14 * c.radius * c.radius
}

// func main() {
// 	c := Circle3{radius: 5}
// 	c.calcArea3()
// 	fmt.Printf("%+v", c)
// }

// Method Sets

type Student2 struct {
	name   string
	grades []int
}

func (s *Student2) displayName() {
	fmt.Println(s.name)
}

func (s *Student2) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

// func main() {
// 	s := Student2{name: "Joe", grades: []int{90, 75, 80}}
// 	s.displayName()
// 	fmt.Printf("%.2f%%", s.calculatePercentage())
// }

type Movie struct {
	name    string
	summary string
	rating  float32
}

func (m Movie) getSummary() {
	m.summary = "summary"
}

func (m *Movie) increaseRating() {
	m.rating += 1
}

// func main() {
// 	mov := Movie{"xyz", "", 2.1}
// 	fmt.Printf("%+v", mov)
// 	mov.increaseRating()
// 	mov.getSummary()
// 	fmt.Printf("%+v", mov)
// }

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// }

type Employee struct {
	eid int
	id  int
}

func (e Employee) get_id() int {
	return e.eid + 10
}

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 		employees[i] = Employee{eid: i}
// 		employees[i].id = employees[i].get_id()
// 		fmt.Printf("%+v\n", employees[i])
// 	}
// }

// Interfaces

type Shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printData(s Shape) {
	fmt.Println("Shape: ", s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
	fmt.Println()
}

func main() {
	r := rect{length: 3, breadth: 4}
	c := square{side: 5}
	printData(r)
	printData(c)
}
