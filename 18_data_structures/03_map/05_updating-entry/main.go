package main

import "fmt"

func main() {
	myMap := map[string]string{
		"Bob": "Yo",
	}

	myMap["Bob"] = "Sup"

	fmt.Println(myMap) // now "Sup"
}
