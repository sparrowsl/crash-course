package main

import "fmt"

func main() {
	guestList := []string{"Melkey", "Johnny Depp", "Linus Trovalds"}

	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[0])
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[1])
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[2])
}
