package main

import "fmt"

//Contact exported
type Contact struct {
	greeting string
	name     string
}

//SwitchOnType exported
func SwitchOnType(x interface{}) {
	switch x.(type) { // type assertion(type is a lexical element keyword)
	case int:
		fmt.Println("Sup Integer")
	case string:
		fmt.Println("Sup String")
	case Contact:
		fmt.Println("Sup Contact")
	default:
		fmt.Println("No matches")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("UFO")
	var t = Contact{"Good to see you,", "Jim"}
	fmt.Println("t: ", t)
	SwitchOnType(t)
}
