/* Coding challenge #5: Calculate the sum of numbers from 1 to 10*/
package main

import (
	"fmt"
)

func main() {
	var acc int = 0
	for i := 1; i <= 10; i++ {
		acc += i
	}
	fmt.Println(acc)
}
