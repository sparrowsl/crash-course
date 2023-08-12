package main

import (
	"fmt"
	"strings"
)

func main() {
	currentUsers := []string{"john", "jane", "adam", "jack", "jenny", "barry", "joe"}
	newUsers := []string{"jonas", "jane", "Jenny", "angela", "jack"}

	for _, name := range newUsers {
		if contains(currentUsers, strings.ToLower(name)) {
			fmt.Println("Username already taken, choose another name.")
		} else {
			fmt.Println("Username is available!")
		}
	}
}

func contains(slice []string, item string) bool {
	inSlice := false
	for _, v := range slice {
		if item == v {
			inSlice = true
		}
	}
	return inSlice

}
