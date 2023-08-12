package main

import "fmt"

func main() {
	numbers := []int{}

	// Make a list of number from 1 to 1,000,000
	for i := 1; i <= 1_000_000; i++ {
		numbers = append(numbers, i)
	}

	// Print each number
	for i := 1; i <= len(numbers); i++ {
		fmt.Println(i)
	}
}
