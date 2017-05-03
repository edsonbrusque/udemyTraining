// The usual way to write a function
package main

import "fmt"

func half(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true
	} else {
		return n / 2, false
	}
}

func main() {
	var n int
	var b bool
	n, b = half(2)
	fmt.Println(n, b)
}
