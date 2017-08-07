package main

import (
	"fmt"
)

// make sets length and capacity

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
