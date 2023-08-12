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

type IceCreamStand struct {
	restaurant Restaurant
	flavors    []string
}

func (i IceCreamStand) showFlavors() {
	for _, flavor := range i.flavors {
		fmt.Println(flavor)
	}
}

func main() {
	ics := IceCreamStand{flavors: []string{"vanilla", "chocolate", "cream"}}
	ics.showFlavors()
}
