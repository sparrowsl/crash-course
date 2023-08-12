package main

import "fmt"

func main() {
	ordinalNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range ordinalNumbers {
		if num == 1 {
			fmt.Println("1st")
		} else if num == 2 {
			fmt.Println("2nd")
		} else if num == 3 {
			fmt.Println("3rd")
		} else {
			fmt.Printf("%dth\n", num)
		}
	}
}
