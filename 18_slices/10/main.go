package main

import "fmt"

func main() {
	{
		var student []string
		var students [][]string

		//student[0] = "Todd"
		student = append(student, "Todd")

		fmt.Println()
		fmt.Println(student == nil, student)
		fmt.Println(students == nil, students)
	}

	{
		student := []string{}
		students := [][]string{}

		//student[0] = "Todd"
		student = append(student, "Todd")

		fmt.Println()
		fmt.Println(student == nil, student)
		fmt.Println(students == nil, students)
	}

	{
		student := make([]string, 35)
		students := make([][]string, 35)

		student[0] = "Todd"
		student = append(student, "Todd")

		fmt.Println()
		fmt.Println(student == nil, student)
		fmt.Println(students == nil, students)
	}

}
