/*Coding challenge #4: Print all the multiplication tables with numbers from 1 to 10*/

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "X", j, "=", i*j)
		}
	}
}
