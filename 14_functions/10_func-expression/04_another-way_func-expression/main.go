package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world 2!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
