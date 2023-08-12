package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("What kind of Car would you like? ")
	car, _ := scanner.ReadString('\n')

	fmt.Printf("Let me see if I can find you a %s", car)
}
