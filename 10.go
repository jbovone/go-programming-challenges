/*Coding challenge #10: Calculate the sum of numbers in an array of numbers*/
package main

import (
	"fmt"
)

func main() {
	numbers := [5]int{2, 2, 5, 4, 5}
	var acc int = 0 //18
	for index, value := range numbers {
		acc += value
		fmt.Println(index)
	}
	fmt.Println(acc)

}

/*
C:\Users\PC\go\src\go>go run 10.go
# command-line-arguments
.\10.go:11:6: index declared but not used
*/
