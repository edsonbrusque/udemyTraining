package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z square) perimeter() float64 {
	return 4 * z.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type triangle struct {
	side float64
}

func (t triangle) area() float64 {
	return t.side * t.side * math.Sqrt(3) / 4
}

func (t triangle) perimeter() float64 {
	return 3 * t.side
}

type shape interface {
	area() float64
	perimeter() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Println(z.perimeter())
}

func main() {
	s := square{10}
	c := circle{5}
	t := triangle{20}
	info(s)
	info(c)
	info(t)
}
