package main

import "fmt"

func main() {
	people := map[string]int{
		"Jack":  20,
		"Jane":  4,
		"Joe":   10,
		"Jill":  7,
		"Jenny": 2,
	}

	for name, age := range people {
		fmt.Printf("%s's favorite number is %d\n", name, age)
	}
}
