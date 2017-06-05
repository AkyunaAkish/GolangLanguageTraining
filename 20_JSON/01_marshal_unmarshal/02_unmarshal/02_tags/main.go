package main

import (
	"encoding/json"
	"fmt"
)

// Person exported
type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"` // json is Unmarshalled with a field of "wisdom score" store that value as the Age field
}

func main() {
	var p1 Person
	fmt.Println(p1.First) // empty
	fmt.Println(p1.Last)  // empty
	fmt.Println(p1.Age)   // 0

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":"20"}`)
	json.Unmarshal(bs, &p1) // Unmarshal takes a slice of bytes and the address of a struct(in this case at least)
	fmt.Println(p1)         // {James Bond 0}
	fmt.Printf("%T \n", p1) // main.Person
}
