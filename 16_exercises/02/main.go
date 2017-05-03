// Now, using a func expression
package main

import "fmt"

func main() {

	half := func(n int) (int, bool) {
		if n%2 == 0 {
			return n / 2, true
		} else {
			return n / 2, false
		}
	}

	var n int
	var b bool

	n, b = half(1)
	fmt.Println(n, b)

	n, b = half(2)
	fmt.Println(n, b)
}
