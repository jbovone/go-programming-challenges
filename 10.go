/*Coding challenge #10: Calculate the sum of numbers in an array of numbers*/
package main

import "fmt"

func main() {
	myArr := []int{2, 3, 4}
	fmt.Println(
		sumArr(myArr),
	)
}
func sumArr(arr []int) int {
	sumArr := 0
	for value := range arr {
		sumArr += arr[value]
	}
	return sumArr
}

/*
C:\Users\PC\go\src\go>go run 10.go
# command-line-arguments
.\10.go:11:6: index declared but not used
*/
