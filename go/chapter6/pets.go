package main

import "fmt"

func main() {
	pets := map[string]map[string]interface{}{
		"sparrow": {
			"animal": "Bird",
			"owner":      "Sparrow",
		},
		"bolt": {
			"animal": "Dog",
			"owner":      "Sparrow",
		},
	}

	for key, value := range pets {
		fmt.Printf("%s\n", key)
		for k, v := range value {
			fmt.Printf("%s: %v\n", k, v)
		}
		fmt.Println()
	}
}
