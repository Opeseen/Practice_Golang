package main

// Create an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Create a rectangle struct
type Rectangle struct {
	Length, Width float64
}

// Create an area method
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Create a perimeter method
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// func main() {
// 	var s Shape = Rectangle{4.0, 6.0}
// 	fmt.Println(s.Area())
// 	fmt.Println(s.Perimeter())
// }

// -----------------------------------------------

// func main() {
// 	r := strings.NewReader("Learning is fun")
// 	buf := make([]byte, 4)
// 	for {
// 		n, err := r.Read(buf)
// 		fmt.Println(buf[:n], err)
// 		if err != nil {
// 			fmt.Println("Breaking out")
// 			break
// 		}
// 	}
// }

// func main() {
// 	r := strings.NewReader("Learning is fun")
// 	buf := make([]byte, 4)
// 	for {
// 		n, err := r.Read(buf)
// 		fmt.Println(string(buf[:n]), err)
// 		if err != nil {
// 			fmt.Println("Breaking out")
// 			break
// 		}
// 	}
// }

// func main() {
// 	r := strings.NewReader("Some io.reader stream to be read\n")
// 	if _, err := io.Copy(os.Stdout, r); err != nil {
// 		log.Fatal(err)
// 	}
// }
