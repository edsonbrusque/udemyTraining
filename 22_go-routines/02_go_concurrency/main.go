package main

import "fmt"

func main() {
	go foo()
	go bar()
	// Nothing happens because main() exits before anything gets printed
}

func foo() {
	for i := 0; i < 1000000; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 1000000; i++ {
		fmt.Println("Bar:", i)
	}
}
