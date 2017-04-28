package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	//increment is set to the return value of wrapper
	//which is another function
	//the inner returned function originally had scope
	//to "x"
	//so the language will hold onto the execution context
	//that the returned function was within in order to
	//preserve the value/reference of/to "x"
	//therefore calling increment() will work and
	//"x" will increment by 1 each time it's called
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
