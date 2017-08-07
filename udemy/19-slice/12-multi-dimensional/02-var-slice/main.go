package main

import (
	"fmt"
)

// var method will always set whatever you are declaring to the zero value

func main() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
