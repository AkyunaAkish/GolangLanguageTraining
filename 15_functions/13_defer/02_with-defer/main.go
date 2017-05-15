package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Print("world")
}

func main() {
	defer world()
	// defer pushes this function call
	// to the end of the callstack (right
	// before the main function exits)
	hello()
}

//defer is useful for closing a file because
//it happens at the end of the main func
//closing files with src, err := os.Open("src.txt")
//defer src.Close()
