package main

import "fmt"

var x = "Hello, World"
var y = "Bye"

// this will not work
// x := "Hello, world"

func main() {
	fmt.Println(x)
}

func f() {
	fmt.Println(y)
}
