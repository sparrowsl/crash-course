package main

import (
	"fmt"
	"strings"
)

func main() {
	filename := "golang_notes.txt"
	fmt.Println(strings.TrimSuffix(filename, ".txt"))
}
