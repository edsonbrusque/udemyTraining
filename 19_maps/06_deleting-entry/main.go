package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":     "Good morning!",
		"Jenny":   "Bonjour!",
		"Harleen": "Howdy!",
	}

	fmt.Println(myGreeting)

	delete(myGreeting, "nobody")
	fmt.Println(myGreeting)

	delete(myGreeting, "Jenny")
	fmt.Println(myGreeting)
}
