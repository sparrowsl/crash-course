package main

import "fmt"

func main() {
	names := []string{"john", "jane", "admin", "jack", "jenny"}
	names = []string{}

	if len(names) != 0 {
		for _, name := range names {
			if name == "admin" {
				fmt.Println("Hello admin, would you like to see a status report?")
			} else {
				fmt.Printf("Hello %s, thank you for logging in again.\n", name)
			}
		}
	} else {
		fmt.Println("We need to find some users!!")
	}
}
