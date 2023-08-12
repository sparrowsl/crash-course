package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "\tJackie\n"
	fmt.Println(name)
	fmt.Println(strings.TrimLeft(name, "\t"))
	fmt.Println(strings.TrimRight(name, "\n"))
	fmt.Println(strings.TrimSpace(name))
}
