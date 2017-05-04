package main

import "fmt"

func main() {
	var x [64]string

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < len(x); i++ {
		x[i] = string(i + 64)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
