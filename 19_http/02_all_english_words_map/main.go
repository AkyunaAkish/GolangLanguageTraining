package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)
	//bufio is an input output buffer
	//a buffer is a way of storing a lot of data
	//at once and then that buffer can be read all
	//at the same time instead of receiving one
	//piece of data at a time
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	defer res.Body.Close()

	//this goes over the scanned array
	for sc.Scan() {
		//this sets a key/value pair in words
		//of the scanned text value for the current item
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}

	for k := range words {
		fmt.Println(k)
	}
}
