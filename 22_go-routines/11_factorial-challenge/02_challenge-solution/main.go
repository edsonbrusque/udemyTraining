package main

import "fmt"

func main() {
	c := chan int
	
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1
-- Use goroutines and channels to calculate factorial

CHALLENGE #2
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your answer, then post that answer to this discussion here: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/

