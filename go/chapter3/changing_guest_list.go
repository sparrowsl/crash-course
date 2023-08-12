package main

import "fmt"

func main() {
	guestList := []string{"Melkey", "Johnny Depp", "Linus Trovalds"}

	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[0])
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[1])
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[2])

	fmt.Printf("\nUnfortunately %s can't make it.\n\n", guestList[2])

	guestList = []string{"Melkey", "Johnny Depp", "Rich Harris"}
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[0])
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[1])
	fmt.Printf("Dear %s, you are invited to my party!!\n", guestList[2])
}
