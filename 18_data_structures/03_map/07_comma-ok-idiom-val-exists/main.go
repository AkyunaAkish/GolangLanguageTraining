package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Bob":      "Good Morning",
		"Vladamir": "Dobroe Ootro",
	}

	//the following if statement
	//gets val and exist and scopes those
	//values to within the if/else block of code
	//then after the semicolon, the exists value is
	//being checked as a boolean
	if val, exists := myGreeting["Bob"]; exists {
		fmt.Println("That val does exist ", val, exists)
	} else {
		fmt.Println("That val does not exist ", val, exists)
	}
}
