package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Buenos dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
		"Bom dia!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	fmt.Println(greeting)

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("[1:2] -> ", greeting[1:2])
	fmt.Println("[:2] --> ", greeting[:2])
	fmt.Println("[5:] --> ", greeting[5:])
	fmt.Println("[:] ---> ", greeting[:])
	fmt.Println("--------------------------------------------------------------------------------")
}
