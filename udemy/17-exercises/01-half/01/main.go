package main

import "fmt"

func halfEven(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	y, even := halfEven(17)
	fmt.Println(halfEven(10))
	fmt.Println(halfEven(5))
	fmt.Println(y, even)
}
