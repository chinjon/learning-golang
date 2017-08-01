package main

import "fmt"

func main() {
	halfEven := func(x int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}

	half, even := halfEven(12)
	fmt.Println(half, even)
}
