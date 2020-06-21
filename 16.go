/*Coding challenge #16: Create a function that will return a Boolean specifying if a number is prime*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPrime(6))
}

func isPrime(number int) bool {
	prime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			prime = false
			break
		}
	}
	return prime
}
