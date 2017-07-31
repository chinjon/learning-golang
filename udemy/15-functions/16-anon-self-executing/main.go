package main

import "fmt"

// anonymous self-executing function
func main() {
	func() {
		fmt.Println("I'm driving!")
	}() // parenthesis execute function after declaring it
}
