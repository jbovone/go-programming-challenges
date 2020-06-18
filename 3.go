/*Coding challenge #3: Print the multiplication table with 7*/

package main

import (
	"fmt"
)

func main() {
	const seven int = 7

	for i := 1; i <= 10; i++ {
		fmt.Println(seven, "X", i, "=", seven*i)
	}
}

/*
	for i := 1; i <= 10; i++ {
		const result string = seven + "X" + i + "=" + seven*i
		fmt.Println(result)
	}
C:\Users\PC\go\src\go>go run 3.go
# command-line-arguments
.\3.go:11:31: cannot use "X" (type untyped string) as type int
*/

/*
 	const seven int = 7
	const X string = "X"

	for i := 1; i <= 10; i++ {
		fmt.Println(X + seven)
	}
	# command-line-arguments
.\3.go:14:17: invalid operation: X + seven (mismatched types string and int)
	fmt.Println(X, seven)
*/
