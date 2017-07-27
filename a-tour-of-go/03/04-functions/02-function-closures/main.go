package main

import "fmt"

// GO functions may be closures.
// a closure is a function value that references variables from outside the its body.

// the function may access and assign to the referenced variables from outside its body.

// the function may access and assign to the referenced variables - in this sense the function is "bound" to the variables

func adder() func(int) int {
	// adder() returns a closure
	// each closure is bound to its own sum variable
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
