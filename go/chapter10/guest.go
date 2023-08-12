package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter name: ")
	input.Scan()

	file, err := os.OpenFile("./guest.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(input.Text())
	if err != nil {
		fmt.Println("Error writing to file")
		panic(err)
	}
}
