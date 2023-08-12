package main

import "fmt"

func main() {
	pizzas := []string{"pepper", "pizza", "phone", "cheese", "vanilla", "laptop", "watch"}

	fmt.Println("The first three items in the list are:")
	fmt.Printf("\t%v\n", pizzas[:3])

	fmt.Println("Three items from the middle of the list are:")
	fmt.Printf("\t%v\n", pizzas[len(pizzas)/2:len(pizzas)/2+3])

	fmt.Println("The last three items in the list are:")
	fmt.Printf("\t%v\n", pizzas[len(pizzas)-3:])
}
