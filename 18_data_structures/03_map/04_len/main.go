package main

import "fmt"

func main() {
	myMap := map[string]string{
		"Bob": "Yo",
	}

	myMap["Joe"] = "Sup"

	fmt.Println(len(myMap))
}
