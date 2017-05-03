// (true && false) || (false && true) || !(false && false)
// false || false || true
// true

package main

import "fmt"

func foo(i ...int) {
	fmt.Print("[")
	for _, j := range i{
		fmt.Print(j, " ")
	}
	fmt.Println("]")
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
