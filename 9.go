package main

import (
	"fmt"
)

func main() {
	fmt.Println(celcToFah(68))
}
func celcToFah(n float64) float64 {
	var result float64 = (n - 32) / 1.8
	return result
}
