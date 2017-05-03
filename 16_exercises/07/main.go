// Project Euler
// Solution to problem 2
// https://projecteuler.net/problem=2

package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 0
	sum := 2
	for {
		c = a + b

		if c > 4000000 {
			break
		}

		if c%2 == 0 {
			sum += c
		}

		a = b
		b = c
	}
	fmt.Print("The sum is: ", sum)
}
