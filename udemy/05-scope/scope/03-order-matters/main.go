package main

import "fmt"

func main() {
	fmt.Println(x) // this will not run, x is declared after
	fmt.Println(y) // will run, y is within the global scope
	x := 42
}

var y = 42
