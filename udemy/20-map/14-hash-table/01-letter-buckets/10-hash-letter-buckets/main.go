package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)

	for scanner.Scan() {
		// assigns that index within the slice the value
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])

	fmt.Println("******************")
	for i := 28; i <= 126; i++ {
		fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	}
}

func hashBucket(word string) int {
	// gets a number based on the first letter of the word
	return int(word[0])
}
