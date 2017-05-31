package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// doesn't work
// moby dick returns a lot of data
// and a panic index out of range error occurs
func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")

	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	// close res body right before function ends/exits
	// res.Body is a reader type
	// and will stay open and keep memory space if you don't close it
	defer res.Body.Close()

	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, 200)

	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	// fmt.Println(buckets[65:123])

	// each symbol/character has a position in the array which
	// its ASCII number corresponds to
	// each time the loop scanner.Scan() ran the value for each
	// character incremented
	fmt.Println(buckets)
}

// HashBucket exported func
// which converts the first letter of each word to its
// ASCII number equivalent
func HashBucket(word string) int {
	return int(word[0])
}
