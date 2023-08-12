package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {

	displayUserInfo()
	// saveUserInfo()
}

func saveUserInfo() {
	scanner := bufio.NewScanner(os.Stdin)
	user := struct {
		name    string
		age     int
		contact string
	}{}

	fmt.Printf("Enter name: ")
	scanner.Scan()
	user.name = scanner.Text()

	fmt.Printf("Enter age: ")
	scanner.Scan()
	intAge, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("Please enter a valid age!!")
		return
	}
	user.age = int(intAge)

	fmt.Printf("Enter contact: ")
	scanner.Scan()
	user.contact = scanner.Text()

	userBytes, _ := json.Marshal(map[string]interface{}{"name": user.name, "age": user.age, "contact": user.contact})
	os.WriteFile("./remember.json", userBytes, os.ModePerm)
}

func displayUserInfo() {
	userFile, err := os.ReadFile("./remember.json")
	if err != nil {
		fmt.Println("File does not exist!!")
		return
	}

	details := map[string]interface{}{}
	json.Unmarshal(userFile, &details)
	fmt.Println(details)
}
