package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

// HashBucket exported
func HashBucket(word string, buckets int) int {
	letter := int(word[0])     //ASCII number for G
	bucket := letter % buckets //remainder divided by 12
	return bucket              //11
}

// HashTables are the data structure that maps use under the hood
// instead of search an unsorted array/slice, or a sorted array/slice
// HashTables create similar length/size buckets for data that
// is similar, in doing so, searching for items is much faster
// if you have a bucket for strings that start with the letter "S"
// then you can go to the "S" bucket and search through that subset
// of data etc...
// this function is supposed to show how the word
// "Go" could go into a bucket 11 because the first letter's
// ASCII number divided by 12 is 11 and that can be used
// to determine which bucket it belongs to
