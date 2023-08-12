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

	fmt.Printf("How many people are in the group? ")
	people, _ := input.ReadString('\n')

	peopleNumber, err := strconv.ParseInt(strings.TrimSpace(people), 10, 64)
	if err != nil {
		fmt.Println("Please enter a valid number!")
		return
	}

	if int(peopleNumber) > 8 {
		fmt.Println("Sorry, but you will have to wait for a table")
	} else {
		fmt.Println("The table is ready!!")
	}
}
