package main

import (
	"encoding/json"
	"fmt"
)

// Person unexported fields
type Person struct {
	first       string
	last        string
	age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1) // turns golang datatype into JSON (to be used in most circumstances except for writing JSON to a stream)
	fmt.Println(bs)           // print byte slice
	fmt.Printf("%T \n", bs)   // []uint8
	fmt.Println(string(bs))   // print JSON
}
