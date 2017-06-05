package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person with almost all exported fields
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

// decoding is used to read JSON from a stream(reader)
func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age": 20}`) // takes a string and gives a pointer to a new reader
	json.NewDecoder(rdr).Decode(&p1)                                        // takes reader and gives pointer to new Decoder then you can use decode method which takes the pointer of a value and will set values in p1 if JSON is valid

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
