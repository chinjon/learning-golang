package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}

// defer is often used when resources need to be freed in some way

/*
	f, _ := os.Open(filename)
	defer f.Close()
*/

// it has three advantages:
// it keeps our Close call near our open call so it's easier to understand
// if our function had multiple return statements, Close will happen before both of them.
// deferred functions are run even if a runtime panic occurs
