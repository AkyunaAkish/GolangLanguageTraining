package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	if err != nil {
		log.Fatal(err)
	}

	// scan page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set split function for scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice of slice of string to hold slices of words
	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// Loop over words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	// print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// print the words in one of the buckets
	fmt.Println(buckets[6])
}

//HashBucket exported
func HashBucket(s string, n int) int {
	var sum int
	for _, v := range s {
		sum += int(v)
	}
	return sum % n
}
