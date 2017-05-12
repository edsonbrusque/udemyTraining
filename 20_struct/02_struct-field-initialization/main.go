package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := Person{"James", "Bond", 20}
	p2 := Person{"Miss", "Moneypenny", 18}
	var p3 Person
	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age)
	fmt.Println(p3.First, p3.Last, p3.Age)
	fmt.Printf("%T %v \n", p1, p1)
	fmt.Printf("%T %v \n", p2, p2)
	fmt.Printf("%T %v \n", p3, p3)

	fmt.Printf("****************************************\n")

	d1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	d2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   18,
		},
		First:         "If Looks Could Kill",
		LicenseToKill: false,
	}

	fmt.Println(d1.First, d1.Person.First)
	fmt.Println(d2.First, d2.Person.First)
}
