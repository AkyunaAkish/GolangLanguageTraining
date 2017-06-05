package main

import (
	"encoding/json"
	"os"
)

// Person with all but one field exported
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

// encoding is used to write JSON to a stream(writer)
func main() {
	p1 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1) // similar affect to Println() in this example
	// NewEncoder takes a writer and gives a pointer to the encoder
	// os.Stdout is a writer, The NewEncoder method then returns
	// the pointer to the writer os.Stdout and which that pointer
	// you can use the  Encode method and pass in a value and turn it into
	// JSON
}
