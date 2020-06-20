/* Coding challenge #13: Find the maximum number in an array of numbers */

package main

import (
	"fmt"
)

func main() {
	myInts := []int{1, 2, 3, 1}
	fmt.Println(
		findMax(myInts),
	)
}
func findMax(arr []int) int {
	maxValue := 0
	for i := range arr {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	return maxValue
}

/*
no an infinity keyword?
 math.Max(...arr) arguments??
*/
