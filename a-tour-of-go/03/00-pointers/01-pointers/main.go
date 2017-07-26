package main

import "fmt"

// * denotes the pointer's underlying value

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println("value for i: ", i, "\tvalue for j: ", j)
	fmt.Println("pointer for i: ", p)
	fmt.Println("reading i through pointer: ", *p)

	*p = 44

	fmt.Println("i now is valued at: ", *p)
	fmt.Println("but still has the same address: ", p)
}
