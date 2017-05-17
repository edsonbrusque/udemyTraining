/*
(2)
s := []string{"Zeno", "John", "Al", "Jenny"}
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string {"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
