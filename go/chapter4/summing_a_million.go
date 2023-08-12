package main

import "fmt"

func main() {
	numbers := []int{}

	for i := 1; i <= 1_000_000; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println(min(numbers))
	fmt.Println(max(numbers))
	fmt.Println(sum(numbers))
}

func min(numbers []int) int {
	min := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}

func max(numbers []int) int {
	max := numbers[0]

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

func sum(numbers []int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}
