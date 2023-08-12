package main

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	sides int
}

func (d Dice) rollDie() {
	fmt.Println(rand.Intn(d.sides) + 1)
}

func main() {
	dice := Dice{sides: 6}
	for i := 1; i <= 10; i++ {
		dice.rollDie()
	}
	dice10 := Dice{sides: 10}
	for i := 1; i <= 10; i++ {
		dice10.rollDie()
	}

	dice20 := Dice{sides: 20}
	for i := 1; i <= 10; i++ {
		dice20.rollDie()
	}
}
