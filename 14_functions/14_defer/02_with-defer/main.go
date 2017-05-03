package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func hello2() {
	fmt.Print("hello2 ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
	hello2()
}
