package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice) // Monday, Tuesday, Wednesday, Thursday, Friday

	// calling the slice up to the value you want and re-append
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice) // Monday, Tuesday, Thursday, Friday
}
