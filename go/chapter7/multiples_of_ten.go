package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a number: ")
	number, _ := input.ReadString('\n')

	// Trim all spaces of the input and convert to number
	numberInt, err := strconv.ParseInt(strings.TrimSpace(number), 10, 64)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	if numberInt%10 != 0 {
		fmt.Printf("%d is not a multiple of 10!\n", numberInt)
	} else {
		fmt.Printf("%d is a multiple of 10!\n", numberInt)
	}
}
