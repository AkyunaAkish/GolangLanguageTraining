package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")

	if err != nil {
		log.Fatal(err)
	}

	pages, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	fmt.Printf("%s \n", pages)
}
