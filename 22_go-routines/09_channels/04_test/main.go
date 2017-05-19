package main

import (
	"fmt"
)

func main() {
	push := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			push <- i
		}
	}()

	pull := make(chan int)
	go func() {
		var sum int
		for n := range push {
			sum += n
		}
		pull <- sum
		close(pull)
	}()

	cSum <- pull(push)
	for n := range cSum {
		fmt.Println(n)
	}
}
