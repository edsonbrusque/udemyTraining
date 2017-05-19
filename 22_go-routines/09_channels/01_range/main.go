package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second/2)
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	fmt.Println("Ended.")
}
