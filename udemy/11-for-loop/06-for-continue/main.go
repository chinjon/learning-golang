package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		// will only print odd numbers
		if i%2 == 0 {
			// if even, it will continue the loop
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
