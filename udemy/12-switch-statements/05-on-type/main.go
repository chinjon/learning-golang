package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Jon")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)

}
