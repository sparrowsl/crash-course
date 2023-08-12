package main

import "fmt"

func main() {
	describeCity("Freetown", "Sierra Leone")
}

func describeCity(city, country string) {
	fmt.Printf("%s is in %s\n", city, country)
}
