package main

import "fmt"

// What is the length of a slice created using make([]int, 3, 9)?

// 3

func main() {
	x := make([]int, 3, 9)

	fmt.Println(x)
}
