package main

import (
	"fmt"
)

type Person struct {
	name string
	edad int
	nac  string
}

func main() {
	x := Person{
		"Jhon Doe",
		24,
		"EEUU",
	}
	fmt.Println(x)
}
