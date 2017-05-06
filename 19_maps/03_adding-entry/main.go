package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy!"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["nobody"])
	fmt.Println(myGreeting["Tim"])
	fmt.Println(myGreeting["Jenny"])
	fmt.Println(myGreeting["Harleen"])
}
