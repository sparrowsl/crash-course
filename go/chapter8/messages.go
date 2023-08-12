package main

import "fmt"

func main() {
	messages := []string{"Hello", "Less is more", "Simple is better than complex"}
	showMessages(messages)
}

func showMessages(messages []string) {
	for _, v := range messages {
		fmt.Println(v)
	}
}
