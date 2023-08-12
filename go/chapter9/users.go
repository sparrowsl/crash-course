package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	age       int
}

func (u User) describeUser() {
	fmt.Printf("Name: %s %s\n", u.firstName, u.lastName)
	fmt.Printf("Age: %d\n", u.age)
}

func (u User) greetUser() {
	fmt.Printf("Hello there %s %s!!\n", u.firstName, u.lastName)
}

func main() {
	john := User{firstName: "John", lastName: "Smith", age: 10}
	jack := User{firstName: "Jack", lastName: "Willborne", age: 20}
	jenny := User{firstName: "Jenny", lastName: "Trush", age: 15}

	john.describeUser()
	john.greetUser()

	jack.describeUser()
	jack.greetUser()

	jenny.describeUser()
	jenny.greetUser()
}
