/*Coding challenge #8: Create a function that will convert from Celsius to Fahrenheit */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(celcToFah(20))
}
func celcToFah(n float64) float64 {
	var result float64 = n*1.8 + 32
	return result
}

/*
func main() {
	fmt.Println(celcToFah(20))
}
func celcToFah(n) {
	const result int = n*1.8 + 32
	return result
}
 command-line-arguments
.\8.go:10:23: celcToFah(20) used as value
.\8.go:12:16: undefined: n
.\8.go:13:2: too many arguments to return
.\8.go:13:9: undefined: n

.\8.go:13:8: const initializer n * 1.8 + 32 is not a constant
.\8.go:13:22: constant 1.8 truncated to integer
*/
