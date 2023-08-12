package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Enter the age: ")
		scanner.Scan()

		age, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Printf("Enter a valid age!!\n\n")
			continue
		}

		switch {
		case age < 3:
			fmt.Println("Movie ticket is free!")
		case age >= 3 && age <= 12:
			fmt.Println("Movie ticket is $10.")
		case age > 12:
			fmt.Println("Movie ticket is $15.")
		}
	}
}
