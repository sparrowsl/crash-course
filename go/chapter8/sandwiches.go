package main

import "fmt"

func main() {
	sandwichItems("cake", "sugar")
	sandwichItems("mint", "pepper", "cream", "milk")
	sandwichItems("chocolate")
}

func sandwichItems(items ...string) {
	fmt.Printf("The sandwich items:\n")

	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println()
}
