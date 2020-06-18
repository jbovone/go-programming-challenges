/*Coding challenge #2: Print the odd numbers less than 100 */
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i += 2 {
		fmt.Println(i)
	}
}
