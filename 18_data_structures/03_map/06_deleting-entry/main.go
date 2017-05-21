package main

import "fmt"

func main() {
	myMap := map[string]string{
		"Bob":  "Yo",
		"Gary": "Meow",
	}

	delete(myMap, "Bob")

	fmt.Println(myMap) // no more bob :'(
}
