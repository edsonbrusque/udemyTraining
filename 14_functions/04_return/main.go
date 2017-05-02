package main

import "fmt"

func main() {
	s := greet("Jane ", "Doe")
	fmt.Println(s)
}

func greet(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}
