/*
use https://godoc.org/sort to sort the following:

(1)
type people []string
studyGroup := people {"Zeno", "John", "Al", "Jenny"}
*/

package main

import (
	"fmt"
	"sort"
)

type people []string

func (a people) Len() int           { return len(a) }
func (a people) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a people) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
