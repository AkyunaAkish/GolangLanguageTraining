package main

import (
	"encoding/json"
	"fmt"
)

// Person exported
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	fmt.Println(p1.First) // empty
	fmt.Println(p1.Last)  // empty
	fmt.Println(p1.Age)   // 0

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":"20"}`)
	json.Unmarshal(bs, &p1) // Unmarshal takes a slice of bytes and the address of a struct(in this case at least)
	fmt.Println(p1)         // {James Bond 0}
	fmt.Printf("%T \n", p1) // main.Person
}
