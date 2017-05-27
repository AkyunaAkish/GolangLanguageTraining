package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour!",
	}

	// other syntax
	//myGreeting := make(map[int]string)
	//myGreeting[0] = "Good Morning!"
	//Bonjour[0] = "Bonjour!"

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
