package main

import (
	"encoding/json"
	"fmt"
)

// Person exported fields
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1) // turns golang datatype into JSON (to be used in most circumstances except for writing JSON to a stream)
	fmt.Println(string(bs))   // {} empty because fields are unexported in Person struct
}
