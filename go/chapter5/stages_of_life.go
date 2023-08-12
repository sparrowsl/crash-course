package main

import "fmt"

func main() {
	age := 10

	if age < 2 {
		fmt.Println("You are a baby")
	} else if age >= 2 && age < 4 {
		fmt.Println("You are a toddler")
	} else if age >= 4 && age < 13 {
		fmt.Println("You are a kid")
	} else if age >= 13 && age < 20 {
		fmt.Println("You are a teenager")
	} else if age >= 20 && age < 65 {
		fmt.Println("You are a adult")
	} else if age >= 65 {
		fmt.Println("You are a elder")
	}

}
