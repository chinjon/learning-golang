package main

import "fmt"

var x = 0 // the variable that is being influenced

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
