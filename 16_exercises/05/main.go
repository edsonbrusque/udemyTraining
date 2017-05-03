// (true && false) || (false && true) || !(false && false)
// false || false || true
// true

package main

import "fmt"

func foo(i ...int) {
	fmt.Printf("%T [", i)
	for _, j := range i {
		fmt.Printf(" %v ", j)
	}
	fmt.Printf("]\n")
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
