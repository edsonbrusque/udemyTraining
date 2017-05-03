// Now, using a func expression
package main

import "fmt"

func theBigger(n... int) int {
	var bigger int

	for _, a := range n {
		if a > bigger {
			bigger = a
		}
	}

	return bigger
}

func main() {
	fmt.Println(theBigger(1, 5, 3, 8, 7, 9, 12, 4))
}
