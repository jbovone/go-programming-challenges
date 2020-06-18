/*Coding challenge #7: Calculate the sum of odd numbers greater than 10 and less than 30 */
package main

import "fmt"

func main() {
	var acc int = 0
	for i := 11; i < 30; i += 2 {
		acc += i
	}
	fmt.Println(acc)
}
