package main

import "fmt"

func main() {
	people := map[string]map[string]interface{}{
		"Jack": {
			"firstName": "Jack",
			"lastName":  "Doe",
			"age":       10,
			"city":      "Freetown",
		},
		"Kate": {
			"firstName": "Kathrine",
			"lastName":  "Smith",
			"age":       13,
			"city":      "Detroit",
		},
	}

	for key, value := range people {
		fmt.Printf("%s\n", key)
		for k, v := range value {
			fmt.Printf("%s: %v\n", k, v)
		}
		fmt.Println()
	}
}
