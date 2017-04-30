package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Print(i, " ")
		if i%2 == 1 {
			fmt.Print("Odd")
		} else {
			fmt.Print("Even")
		}
		fmt.Print("\t")
		if i%10 == 0 {
			fmt.Println()
		}
	}
}
