package main

import "fmt"

const (
	//A exported
	A = iota // 0
	//B exported
	B = iota // 1
	//C exported
	C = iota // 2
)
const (
	//D exported
	D = iota // 0
	//E exported
	E // 1
	//F exported
	F // 2
)

//iota's increment as they are used
func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
