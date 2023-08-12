package main

import "fmt"

func main() {
	makeShirt(10, "Less is More")
}

func makeShirt(size int, text string) {
	fmt.Printf("Shirt size: %d\nShirt message: %s\n", size, text)
}
