package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	//%T is the type of the value
	//%v is a value in the default format
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%v \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%v \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%v \n", d)
}
