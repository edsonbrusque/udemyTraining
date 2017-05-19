package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	semaphore := 2

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		semaphore--
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		semaphore--
	}()

	go func() {
		for semaphore > 0 {
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
