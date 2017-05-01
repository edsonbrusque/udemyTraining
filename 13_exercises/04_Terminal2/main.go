package main

import "fmt"

func main() {
	fmt.Print("Please, enter a small number: ")
	var small int
	fmt.Scan(&small)

	fmt.Print("Please, enter a big number: ")
	var big int
	fmt.Scan(&big)

	fmt.Println(big, " / ", small, " = ", big / small)
}
