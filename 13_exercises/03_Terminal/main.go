package main

import "fmt"

func main() {
	fmt.Print("Please, enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello [", name, "]")
}
