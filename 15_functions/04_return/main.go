package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

//func structure
//func keyword
//reciever
//func name
//func params
//func return type/s
//func body
func greet(fname string, lname string) string {
	//fmt.Sprint is String print, instead of printing
	//to the terminal/console - it concatonates and returns
	//a string
	return fmt.Sprint(fname, " ", lname)
}
