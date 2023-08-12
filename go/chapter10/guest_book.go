package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Enter name: ")
		input.Scan()

		if input.Text() == "quit" {
			break
		}

		file, err := os.OpenFile("./guest_book.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("%s\n", input.Text()))
		if err != nil {
			panic(err)
		}

	}
}
