package main

import (
	"encoding/json"
	"fmt"
)

// Person exported
type Person struct {
	First string
	Last  string `json:"-"`            // this field is not exported when used in json
	Age   int    `json:"wisdom score"` // changes name of field when converted to JSON
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println("Before json marshal: ", p1)
	bs, _ := json.Marshal(p1)
	fmt.Println("After json marshal: ", string(bs))
}
