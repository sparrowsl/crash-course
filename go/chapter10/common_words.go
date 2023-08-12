package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("./romeo_juliet.txt")
	if err != nil {
		panic(err)
	}

	count := strings.Count(strings.ToLower(string(fileBytes)), strings.ToLower("Juliet"))
	fmt.Println(count)
}
