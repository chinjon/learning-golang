// Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and half(2 should return (1, true).

package main

import "fmt"

func half(input int) (int, bool) {
	halved := input / 2
	return halved, true
}

func main() {
	fmt.Println(half(2))
	fmt.Println(half(43))
}
