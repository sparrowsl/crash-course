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

	fmt.Printf("Enter 2 numbers: ")
	numbers, _ := input.ReadString('\n')
	numbersArray := strings.Split(numbers, " ")

	num1, err := strconv.ParseInt(numbersArray[0], 10, 64)
	if err != nil {
		fmt.Println("Please enter a valid number!!")
		return
	}
	num2, err := strconv.ParseInt(numbersArray[1], 10, 64)
	if err != nil {
		fmt.Println("Please enter a valid number!!")
		return
	}

	fmt.Println(num1 + num2)

}
