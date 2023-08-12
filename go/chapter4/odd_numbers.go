package main

import "fmt"

func main() {
	oddNumbers := []int{}

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			oddNumbers = append(oddNumbers, i)
		}
	}

	for _, i := range oddNumbers {
		fmt.Println(i)
	}
}
