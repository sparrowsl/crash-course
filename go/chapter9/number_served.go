package main

import "fmt"

type Restaurant struct {
	name         string
	cuisineType  string
	numberServed int
}

func (res Restaurant) describeRestaurant() {
	fmt.Printf("Name: %s\n", res.name)
	fmt.Printf("Cuisine Type: %s\n", res.cuisineType)
}

func (res Restaurant) openRestaurant() {
	fmt.Printf("The %s Restaurant is now open!\n", res.name)
}

func (res *Restaurant) setNumberServed(num int) {
	res.numberServed = num
}

func (res *Restaurant) incrementNumberServed(num int) {
	res.numberServed += num
}

func main() {
	restaurant := Restaurant{name: "Sunny", cuisineType: "Classic"}

	fmt.Printf("The %s has served %d customers.\n", restaurant.name, restaurant.numberServed)
	restaurant.numberServed = 3
	fmt.Printf("The %s has served %d customers.\n", restaurant.name, restaurant.numberServed)
	restaurant.setNumberServed(10)
	fmt.Printf("The %s has served %d customers.\n", restaurant.name, restaurant.numberServed)
	restaurant.incrementNumberServed(5)
	fmt.Printf("The %s has served %d customers.\n", restaurant.name, restaurant.numberServed)
}
