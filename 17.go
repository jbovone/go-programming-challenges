/*Coding challenge #17: Calculate the sum of digits of a positive integer number*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(
		sumDigits(2354),
	)
}
func sumDigits(number int64) int {
	fmt.Println(number)
	runes := []rune(strconv.FormatInt(number, 10))
	fmt.Println(runes)
	return 2
}

//runes := []rune(strconv.FormatInt(number, 10))
//expect: [2,3,5,4]
//recived: [50 51 53 52]
