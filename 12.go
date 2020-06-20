/* Coding challenge #12: Create a function that receives an array of numbers and returns an array containing only the positive numbers*/
package main

import (
	"fmt"
)

func main() {
	myArr := []int{-1, 2, 4, -5}
	fmt.Println(
		filterNegatives(myArr),
	)
}
func filterNegatives(arr []int) []int {
	arrPositives := []int{}
	for value := range arr {
		if arr[value] > 0 {
			arrPositives = append(arrPositives, arr[value])
		}
	}
	return arrPositives
}
