package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (z Square) perimeter() float64 {
	return 4 * z.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Triangle struct {
	side float64
}

func (t Triangle) area() float64 {
	return t.side * t.side * math.Sqrt(3) / 4
}

func (z Triangle) perimeter() float64 {
	return 3 * z.side
}

type Shape interface {
	area() float64
	perimeter() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Println(z.perimeter())
}

func main() {
	s := Square{10}
	c := Circle{5}
	t := Triangle{20}
	info(s)
	info(c)
	info(t)
}
