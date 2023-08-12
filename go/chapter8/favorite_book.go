package main

import "fmt"

func main() {
	favoriteBook("Harry Potter")
	favoriteBook("Narnia")
}

func favoriteBook(title string) {
	fmt.Printf("One of my favorite books is %s.\n", title)
}
