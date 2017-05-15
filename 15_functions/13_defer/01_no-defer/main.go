package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Print("world")
}

func main() {
	world()
	hello()
}
