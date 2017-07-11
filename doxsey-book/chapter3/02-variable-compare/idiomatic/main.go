package main

import "fmt"

func main() {
	x := "hello"
	y := "world"
	fmt.Println(x == y) // should print false

	// type is not necessary because the Go compiler is able to infer the type based on the literal value you assign the variable.
}
