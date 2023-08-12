package main

import "fmt"

func main() {
	glossary := map[string]string{
		"variable": "A container to store or hold a value",
		"golang":   "A weird but cool language to work with",
		"TDD":      "Test Driven Development, write code to test your code",
		"frontend": "What the users see and interract with",
		"svelte":   "The best javascript framework currently",
	}

	for key, value := range glossary {
		fmt.Printf("%s:\n\t%s\n\n", key, value)
	}
}
