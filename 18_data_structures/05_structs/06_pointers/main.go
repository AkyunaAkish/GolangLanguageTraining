package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20} // assigns the address to p1
	fmt.Println(p1)
	fmt.Println(&p1) // gives actual pointer hexadecimal value
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	// fmt.Println(*p1.name) golang is adding the "*" for you which gives you the value instead of the pointer
	fmt.Println(p1.age)
	// &{James 20} // this means you have a pointer but here are the values
	// 0xc42006c018 // pointer
	// *main.person // the type is a pointer for type person
	// James // name
	// 20 // age
}
