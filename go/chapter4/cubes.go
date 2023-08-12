package main

import "fmt"

func main() {
	cubes := []int{}

	for i := 1; i <= 10; i++ {
		cubes = append(cubes, i)
	}

	for _, i := range cubes {
		fmt.Println(i * i * i)
	}
}
