package main

import "fmt"

func main() {
	sandwichOrders := []string{"bread", "apple", "vanilla", "brown", "tuna"}
	finishedSandwiches := make([]string, 0)

	for _, v := range sandwichOrders {
		fmt.Printf("I made your %s sandwich\n", v)
		finishedSandwiches = append(finishedSandwiches, v)
	}

	fmt.Println(finishedSandwiches)
}
