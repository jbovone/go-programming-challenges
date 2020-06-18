/*Coding challenge #1: Print numbers from 1 to 10*/
package main

import (
	"fmt"
)

func main() {
	const BREAK int = 10
	for i := 1; i <= BREAK; i++ {
		fmt.Println(i)
	}
}

/*
.\1.go:15:1: syntax error: non-declaration statement outside function body
const BREAK int = 10
for i :=0; i < BREAK; i++ {
	fmt.Println(i)

}
*/
