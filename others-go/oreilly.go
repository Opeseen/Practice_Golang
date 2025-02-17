package main

import "fmt"

func createSlice(n int) []string {
	slice := make([]string, 0, n*7)
	for i := 0; i < n; i++ {
		slice = append(slice, "I", "am", "going", "to", "take", "some", "space", ",")
	}
	return slice
}

func main() {
	// Call the function
	result := createSlice(1)
	fmt.Println(result)
}
