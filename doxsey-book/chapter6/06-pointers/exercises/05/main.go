// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).

package main

import "fmt"

func fib(x int) int {

	if x <= 1 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println(fib(5))
}
