package main

import "fmt"

func main() {
	var c []chan int
	for i := 0; i < 100; i++ {
		c = append(c, factorial(i))
	}

	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have wich helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https:/goo.gl/uJa99G
 */
