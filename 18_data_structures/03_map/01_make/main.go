package main

import "fmt"

func main() {
	//map with a key of string and value of string
	//use make() almost always so you can set values with
	//square bracket notation. Without make the map is nil
	//and won't accept square bracket value setting notation

	//Another way to properly set a map could be
	//var m = map[string]string{} //using curly braces for "composite literal notation"
	var myGreeting = make(map[string]string)
	myGreeting["Bob"] = "Good Morning"
	myGreeting["Vladamir"] = "Dobroe Ootro"

	fmt.Println(myGreeting)
}
