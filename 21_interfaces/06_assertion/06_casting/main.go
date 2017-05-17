package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T %v\n", rem, rem)
	fmt.Printf("%T %v\n", int(rem), rem)

	var val interface{} = 7
	fmt.Printf("%T %v\n", val, val)
	//fmt.Printf("%T %v\n", int(val), val) // cannot convert val (type interface {}) to type int: need type assertion
	fmt.Printf("%T %v\n", val.(int), val) // cannot convert val (type interface {}) to type int: need type assertion
}
