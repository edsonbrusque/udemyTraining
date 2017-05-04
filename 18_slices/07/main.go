package main

import "fmt"

func main() {
	{
		mySlice := []int{1, 2, 3, 4, 5}
		myOtherSlice := []int{6, 7, 8, 9}

		mySlice = append(mySlice, myOtherSlice...)
		fmt.Println(mySlice)
	}

	{
		mySlice := []string{"Monday", "Tuesday"}
		myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

		mySlice = append(mySlice, myOtherSlice...)
		fmt.Println(mySlice)

		mySlice = append(mySlice[:2], mySlice[3:]...)
		fmt.Println(mySlice)
	}
}
