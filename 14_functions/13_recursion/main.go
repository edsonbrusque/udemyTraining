package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func factorial2(x int) int {
	fact := 1
	for ; x > 1; x-- {
		fact *= x
	}
	return fact
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial2(0))
	fmt.Println(factorial2(1))
	fmt.Println(factorial2(2))
	fmt.Println(factorial2(3))
	fmt.Println(factorial2(4))
	fmt.Println(factorial2(10))
}
