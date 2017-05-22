package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		total := 1
		for i := <-in; i > 0; i-- {
			total *= i
		}
		out <- total
	}()

	in <- 4
	fmt.Println("Factorial:", <-out)
}

/*
CHALLENGE #1
-- Use goroutines and channels to calculate factorial

CHALLENGE #2
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your answer, then post that answer to this discussion here: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
