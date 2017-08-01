package main

import "fmt"

func halfEven(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	half, even := halfEven(13)
	fmt.Println(half, even)
}
