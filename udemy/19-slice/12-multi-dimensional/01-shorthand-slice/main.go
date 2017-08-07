package main

import "fmt"

// shorthand method, and slice will not be nil
// need to use append

func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
