package main

import (
	"fmt"
)

func main() {
	print("Please, enter your name:")
	var name string
	fmt.Scanln(&name)
	println("Hello ", name)
}
