package main

import "fmt"

func main() {
	name := "Todd"
	// #1
	fmt.Println(name) // Todd

	// #2
	changeMe(name)

	// #3
	fmt.Println(name)
}

func changeMe(z string) {

	// #2 - #1
	fmt.Println(z) // grabbing from global scope
	z = "Rocky"
	// #2 - #2
	fmt.Println(z) // grabbing from local scope
}
