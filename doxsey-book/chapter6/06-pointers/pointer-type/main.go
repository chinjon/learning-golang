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

// when we write *xPtr = 0, we are saying
// "store the int 0 in the memory location xPtr refers to."
// if we try xPtr = 0 instead, we will get a compile-time error because xPtr is not an int
// it is a *int, which can only be given another *int

// we use the & operator to find the addess of a variable.
// &x returns a *int (pointer to an int) because x is an int. This is what allows us to modify the original variable.
// &x in main and xPtr in zero refer to the same memory location.
