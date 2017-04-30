package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters int32
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := float64(meters) * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
