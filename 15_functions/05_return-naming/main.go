package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	//automatically returns s if s has been assigned
	return
}
