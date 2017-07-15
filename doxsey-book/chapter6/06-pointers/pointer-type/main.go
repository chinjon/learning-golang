package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

// pointers reference a location in memory where a value is stored rather than the value itself.
// by using a pointer (*int), the zero function is able to modify the original variable.

// in Go, a pointer is represented using an asterisk (*) followed by the type of the stored value

// an asterisk is also used to dereference pointer variables.
// dereferencing a pointer gives us access to the value of the pointer points to.
