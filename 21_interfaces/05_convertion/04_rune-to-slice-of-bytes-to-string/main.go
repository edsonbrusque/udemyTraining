package main

import "fmt"

func main() {
	fmt.Println(string([]byte{72, 101, 108, 108, 111}))
	fmt.Println(string([]byte{'H', 'e', 'l', 'l', 'o'}))
	// conversion: []bytes to string
	// we'll learn about []bytes soon
}
