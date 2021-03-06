package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20}
	p2 := person{"Ian", 44}

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name, (*p1).name)
	fmt.Println(p1.age, (*p1).age)

	fmt.Println()

	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
	fmt.Println(p2.name)
	fmt.Println(p2.age)
}
