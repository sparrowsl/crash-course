package main

import "fmt"

type Restaurant struct {
	restaurantName string
	cuisineType    string
}

func (res Restaurant) describeRestaurant() {
	fmt.Printf("Name: %s\n", res.restaurantName)
	fmt.Printf("Cuisine Type: %s\n", res.cuisineType)
}

func (res Restaurant) openRestaurant() {
	fmt.Printf("The %s Restaurant is now open!\n", res.restaurantName)
}

func main() {
	restaurant := Restaurant{restaurantName: "Sunny", cuisineType: "Classic"}

	fmt.Println(restaurant.restaurantName)
	fmt.Println(restaurant.cuisineType)

	restaurant.openRestaurant()
	restaurant.describeRestaurant()
}
