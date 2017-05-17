package main

import "fmt"

func main() {
	var name interface{} = 7
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T [%q]\n", str, str)
	} else {
		fmt.Println("Value is not a string")
	}
}
