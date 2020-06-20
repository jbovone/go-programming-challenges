/*Coding challenge #15: Create a function that will find the nth Fibonacci number using recursion */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		fibonacciRecursive(10),
	)
}

func fibonacciRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

/*
expect 55
*/
/*
love it!
*/
