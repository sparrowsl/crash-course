package main

import "fmt"

type User struct {
	firstName     string
	lastName      string
	age           int
	loginAttempts uint
}

func (u User) describeUser() {
	fmt.Printf("Name: %s %s\n", u.firstName, u.lastName)
	fmt.Printf("Age: %d\n", u.age)
}

func (u *User) incrementLoginAttempts() {
	u.loginAttempts += 1
}

func (u *User) resetLoginAttempts() {
	u.loginAttempts = 0
}

func (u User) greetUser() {
	fmt.Printf("Hello there %s %s!!\n", u.firstName, u.lastName)
}

func main() {
	john := User{firstName: "John", lastName: "Smith", age: 10}

	john.incrementLoginAttempts()
	john.incrementLoginAttempts()
	john.incrementLoginAttempts()
	john.incrementLoginAttempts()
	fmt.Println(john.loginAttempts)

	john.resetLoginAttempts()
	fmt.Println(john.loginAttempts)

}
