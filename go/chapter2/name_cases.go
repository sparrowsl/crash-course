package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Jack Smith"

	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Title(name))
}
