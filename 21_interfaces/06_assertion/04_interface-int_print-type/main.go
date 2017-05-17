package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Printf("%T [%v]\n", val, val)
}
