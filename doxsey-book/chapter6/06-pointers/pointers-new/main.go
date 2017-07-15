package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

// new takes a type as an argument, allocates enough memory to fit a value of that type, and returns a pointer to it.

// in some programming languages, there is a significant difference between using new and &, with great care being needed to eventually delete anything create with new.

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
