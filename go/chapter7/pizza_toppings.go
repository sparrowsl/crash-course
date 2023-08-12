package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Enter a pizza topping: ")
		scanner.Scan()

		if scanner.Text() == "quit" {
			break
		}

		fmt.Printf("I will add the %s on your pizza.\n\n", scanner.Text())
	}
}
