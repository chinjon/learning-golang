package main

import (
	"fmt"
)

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham" // can not add, not properly appended

	fmt.Println(greeting[2])
}
