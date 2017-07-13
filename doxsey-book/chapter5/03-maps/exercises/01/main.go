package main

import "fmt"

// How do you access the fourth element of an array or slice?

func main() {
	x := [10]string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
	}

	fmt.Println(x[3])
}
