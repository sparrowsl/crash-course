package main

import "fmt"

func main() {
	messages := []string{"Hello there", "Readability counts", "Less is more", "Simple is better"}
	sentMessages := make([]string, 0)

	sendMessages(messages, &sentMessages)
	fmt.Println(messages)
	fmt.Println(sentMessages)
}

func sendMessages(messages []string, sent *[]string) {
	for len(messages) > 0 {
		firstMessage := (messages)[0]
		*sent = append(*sent, firstMessage)
		messages = (messages)[1:]
	}
}

func showMessages(messages []string) {
	for _, v := range messages {
		fmt.Println(v)
	}
}
