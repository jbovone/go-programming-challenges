/*Coding challenge #11: Calculate the average of the numbers in an array of numbers*/

package main

import (
	"fmt"
)

func main() {
	myArr := []int{40, 4}
	fmt.Println(
		aveArr(myArr),
	)
}
func aveArr(arr []int) int {
	acc := 0
	for value := range arr {
		acc += arr[value]
	}
	return acc / len(arr)
}
