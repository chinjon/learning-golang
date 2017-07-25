package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) // 0xc42000a310
	fmt.Println(&x)        // 0xc42000a310
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // 0xc42000a2d8
	fmt.Println(&x)        // 0xc42000a2d8
	zero(x)
	fmt.Println(x)
}
