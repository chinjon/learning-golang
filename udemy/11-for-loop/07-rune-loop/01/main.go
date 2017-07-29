package main

import "fmt"

// a rune is a character
// use single quotes for a character ''

// also an alias for int32

func main() {
	// for i := 50; i <= 140; i++ {
	// 	fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	// }
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}
