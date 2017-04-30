package main

import "fmt"

func main() {

	a := 43
	b := 34

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d %T \n", &a, a)
	fmt.Printf("%d %T \n", &b, b)
}
