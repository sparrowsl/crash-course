package main

import "fmt"

func main() {
	cityCountry("Freetown", "Sierra Leone")
	cityCountry("Santiago", "Chile")
	cityCountry("Barcelona", "Spain")
}

func cityCountry(city, country string) {
	fmt.Printf("%s, %s\n", city, country)
}
