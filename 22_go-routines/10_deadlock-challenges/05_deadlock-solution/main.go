package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close (c)
	}()

	for x := range c {
		fmt.Println(x)
	}

}

// Why doest this only print zero?
// And what can you do to print all 0 - 9 numbers?

