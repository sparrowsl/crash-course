package main

import "fmt"

func main() {
	person := map[string]interface{}{
		"firstName": "Jack",
		"lastName":  "Doe",
		"age":       10,
		"city":      "Freetown",
	}

	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
	fmt.Println(person["age"])
	fmt.Println(person["city"])
}
