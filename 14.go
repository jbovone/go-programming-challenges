/*Coding challenge #14: Print the first 10 Fibonacci numbers without recursion*/

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(
			fibonacci(i),
		)
	}
}

func fibonacci(n int) int {
	n0 := 0
	n1 := 1
	if n == n0 {
		return n
	} else if n == n1 {
		return n1
	} else {
		for i := 2; i <= n; i++ {
			if i%2 == 0 {
				n0 += n1
			} else {
				n1 += n0
			}
		}
	}

	if n0 > n1 {
		return n0
	} else {
		return n1
	}
}

/*
expect; 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233,
377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, ...
*/
