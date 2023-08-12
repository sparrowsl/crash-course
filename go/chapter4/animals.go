package main

import "fmt"

func main() {
	animals := []string{"Cheetah", "Shark", "Hawk"}

	for _, animal := range animals {
		fmt.Printf("A %s would probably be a good pet.\n", animal)
	}

	fmt.Println("All these animals are wild animals!!")
}
