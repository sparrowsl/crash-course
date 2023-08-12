package main

import (
	"fmt"
	"sort"
)

func main() {
	locations := []string{"Sychelles", "Iceland", "France", "Romania", "Silicon Valley"}
	fmt.Println(locations)

	sort.Strings(locations) // sort the list in asc order
	fmt.Println(locations)

	sort.Sort(sort.Reverse(sort.StringSlice(locations))) // sort the list in desc order
	fmt.Println(locations)
}
