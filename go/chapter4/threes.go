package main

import "fmt"

func main() {
	threes := []int{}

	for i := 3; i <= 30; i++ {
		if i%3 == 0 {
			threes = append(threes, i)
		}
	}

	for _, i := range threes {
		fmt.Println(i)
	}
}
