// https://projecteuler.net/
// Solution to problem 1
// https://projecteuler.net/problem=1

package main

import "fmt"

func main() {
	sum := 0
	for n := 1; n < 1000; n++ {
		if (n%3 == 0) || (n%5 == 0) {
			sum += n
		}
	}
	fmt.Println(sum)
}
