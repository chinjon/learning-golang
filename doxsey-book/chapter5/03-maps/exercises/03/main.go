package main

import "fmt"

// Given the following array, what would x[2:5] give you?
// x := [6]string{"a","b","c","d","e","f"}

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
	// c,d,e
}
