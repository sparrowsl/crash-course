package main

import "fmt"

func main() {
	pizzas := []string{"pepperoni pizza", "cheese pizza", "vanilla pizza"}

	for _, pizza := range pizzas {
		fmt.Printf("I like %s\n",pizza)
	}

	fmt.Println("I really think I like pizza!")
}
