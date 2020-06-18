/* Coding challenge #6: Calculate 10!*/
package main

import "fmt"

func main() {
	var factorial int = 1
	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}
