package main

import (
	"fmt"
	"os"
)

func main() {
	catsFile, err := os.ReadFile("./cats.txt")
	if err != nil {
		fmt.Println("File does not exist!!", err)
		return
	}

	fmt.Println(string(catsFile))

	dogsFile, err := os.ReadFile("./dogs.txt")
	if err != nil {
		fmt.Println("File does not exist!!", err)
		return
	}

	fmt.Println(string(dogsFile))
}
