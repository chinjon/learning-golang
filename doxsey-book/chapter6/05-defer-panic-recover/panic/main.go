package main

import "fmt"

// panic function is used to cause a runtime error
// we can handle a runtime panic with the built-in recover function.
// recover stops the panic and returns the value that was passed to the call to panic

func main() {

	// 	panic("PANIC")
	// str := recover() // this will never happen
	// fmt.Println(str)

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// a panic generally indicates a programmer error or an exceptional condition that there's no easy way to recover from
