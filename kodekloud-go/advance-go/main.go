package main

import (
	"fmt"
	"strings"
)

func calculate(a int, b int) (results []float64) {
	results = append(results, float64(a+b), float64(a-b), float64(a*b), float64(a/b))
	return results
}

// func main() {
// 	fmt.Println(calculate(20, 10))
// 	fmt.Println(calculate(700, 70))
// }

type Item struct {
	Name  string
	Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
	// your code goes here
	var result []Item
	for _, item := range items {
		if item.Price >= minPrice && item.Price <= maxPrice {
			result = append(result, item)
		}
	}
	return result
}

// func main() {
// 	items := []Item{
// 		{Name: "Apple", Price: 0.5},
// 		{Name: "Banana", Price: 0.25},
// 		{Name: "Orange", Price: 0.75},
// 		{Name: "Pineapple", Price: 1.5},
// 	}

// 	fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
// 	fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
// 	fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
// }

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	// your code goes here
	book.Pages = pages
}

// func main() {
// 	book1 := &Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Pages: 180}
// 	book2 := &Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", Pages: 281}
// 	book3 := &Book{Title: "Pride and Prejudice", Author: "Jane Austen", Pages: 279}

// 	updatePages(book1, 210)
// 	fmt.Println(book1)

// 	updatePages(book2, 250)
// 	fmt.Println(book2)

// 	updatePages(book3, 295)
// 	fmt.Println(book3)

// }

type Expense struct {
	Name   string
	Amount float64
	Date   string
}

func Total(expenses []Expense) float64 {
	var total float64
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total
}

func (e Expense) getName() string {
	return e.Name
}

// func main() {
// 	expenses := []Expense{
// 		{"Grocery", 50.0, "2022-01-01"},
// 		{"Gas", 30.0, "2022-01-02"},
// 		{"Restaurant", 40.0, "2022-01-03"},
// 	}

// 	fmt.Println(Total(expenses))
// 	fmt.Println(expenses[2].getName())
// }

type Product struct {
	name     string
	quantity int
	price    float64
}

func printName(p Product) {
	fmt.Println(p.name)
}

func printQuantity(p Product) {
	fmt.Println(p.quantity)
}

func printPrice(p Product) {
	fmt.Println(p.price)
}

// func main() {
// 	var p Product
// 	p.name = "Chair"
// 	p.quantity = 5
// 	p.price = 700

// 	fmt.Println("Product details:")
// 	printName(p)
// 	defer printPrice(p)
// 	printQuantity(p)
// }

func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}
