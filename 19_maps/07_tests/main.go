package main

import "fmt"

func main() {
	var test1 map[string]string // This doesn't work. You don't have how to append to this map
	//test1["1"] = "one"
	fmt.Println(test1)

	var test2 = make(map[string]string) // This does work.
	test2["2"] = "two"
	fmt.Println(test2)

	test3 := make(map[string]string) // This does work.
	test3["3"] = "three"
	fmt.Println(test3)

	test4 := map[string]string{} // This does work.
	test4["4"] = "four"
	fmt.Println(test4)
}
