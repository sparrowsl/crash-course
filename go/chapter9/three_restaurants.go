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
	res1 := Restaurant{restaurantName: "Sunny", cuisineType: "Classic"}
	res1.openRestaurant()
	res1.describeRestaurant()

	res2 := Restaurant{restaurantName: "Nineties", cuisineType: "VIP"}
	res2.openRestaurant()
	res2.describeRestaurant()

	res3 := Restaurant{restaurantName: "Tech", cuisineType: "Nerds"}
	res3.openRestaurant()
	res3.describeRestaurant()

}
